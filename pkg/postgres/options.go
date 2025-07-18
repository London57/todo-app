package postgres

import "time"

type Option func(*Postgres)

func MaxPoolSize(size int) Option {
	return func(p *Postgres) {
		p.maxPoolSize = size
	}
}

func ConnAttepts(attempts int) Option {
	return func(p *Postgres) {
		p.connAttepts = attempts
	}
}

func ConnTimeout(timeout time.Duration) Option {
	return func(p *Postgres) {
		p.connTimeout = timeout
	}
}
