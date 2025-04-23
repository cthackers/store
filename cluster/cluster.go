package cluster

type Config struct{}

type Cluster struct{}

func New() *Cluster {
	return &Cluster{}
}
