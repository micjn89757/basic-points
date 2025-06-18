package main

import "sync"

type Client struct {
	Count int
}

type Singleton struct {
	client *Client
	metux  sync.Mutex
}

func (s *Singleton) Get() (client *Client, err error) {
	if s.client != nil {
		return s.client, nil
	}
	s.metux.Lock()
	defer s.metux.Unlock()
	// double check
	if s.client == nil {
		s.client, err = &Client{Count: 1}, nil
		if err != nil {
			return nil, err
		}
	}
	return s.client, nil
}

func main() {
	for {
		singleton := &Singleton{}
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			singleton.Get()
		}()
		go func() {
			defer wg.Done()
			client, _ := singleton.Get()
			if client != nil && client.Count != 1 {
				panic("count should be 1")
			}
		}()
		wg.Wait()
	}
}
