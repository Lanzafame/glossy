package conf

type Config map[string]string

func (c Config) Get(key string, def string) string {
	got, ok := c[key]
	if ok {
		def = got
	}
	return def
}
