package services

import (
	"context"
	"time"

	"br.com.cleiton.ratelimiter/internal/storage"
	"github.com/go-redis/redis/v8"
)

type Limiter struct {
	storage     storage.Storage
	fillRate    int // Tokens por segundo
	capacity    int // Capacidade máxima do balde
	timeBlocked int // Tempo em minutos que o ip ficarábloqueado
}

const Prefix_Key_TimeBlocked = "TB-"

func NewLimiter(storage storage.Storage, fillRate, capacity int, timeBlocked int) *Limiter {
	return &Limiter{storage: storage, fillRate: fillRate, capacity: capacity, timeBlocked: timeBlocked}
}

func (l *Limiter) Allow(key string) bool {
	ctx := context.Background()

	// Obter o número atual de tokens
	tokens, err := l.storage.Get(ctx, key)
	if err != nil {
		if err == redis.Nil {
			l.storage.Increment(ctx, key, l.fillRate)
			return true
		}
		return false
	}

	// Verificar se há tokens suficientes
	if tokens <= 0 {
		return false
	}

	// Remover um token
	err = l.storage.Decrement(ctx, key)
	if err != nil {
		// Lidar com erro
		return false
	}

	// Agendar a reposição de tokens
	go func() {
		timeBlocked := l.timeBlocked
		timeBlockedRedis, err := l.storage.Get(ctx, Prefix_Key_TimeBlocked+key)
		if err == nil {
			timeBlocked = timeBlockedRedis
		}
		time.Sleep(time.Second * time.Duration(timeBlocked))
		l.storage.Increment(ctx, key, l.fillRate)
	}()

	return true
}
