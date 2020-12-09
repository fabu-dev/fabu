package short

import (
	"fabu.dev/api/pkg/constant"
	"fabu.dev/api/pkg/db"
)

type Pool struct {
	minLength int64
	length    int64
	short     *Short
}

func NewPool() *Pool {
	return &Pool{
		minLength: 200,
		length:    1000,
		short:     NewShort(),
	}
}

func (p *Pool) Buffer() error {
	currentLength, err := db.Redis.SCard(constant.ShortKeyPool).Result()
	if err != nil {
		return err
	}

	if currentLength > p.minLength {
		return nil
	}

	var i int64
	keySlice := make([]interface{}, 0, p.length-currentLength)
	for i = 0; i < p.length-currentLength; i++ {
		keySlice = append(keySlice, p.short.Create())
	}

	err = db.Redis.SAdd(constant.ShortKeyPool, keySlice...).Err()
	if err != nil {
		return err
	}

	return nil
}

func (p *Pool) GetShortKey() (string, error) {
	key, err := db.Redis.SPop(constant.ShortKeyPool).Result()
	if err != nil {
		return "", err
	}

	return key, err
}
