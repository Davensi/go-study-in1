package discovery



// Register grpc服务器注册到etcd
type Register struct {
	// 服务名称
	ServiceName string
	// 服务地址
	ServiceAddr string
	// 服务端口
	ServicePort int
	// 服务标签
	ServiceTags []string
	// 服务元数据
	ServiceMetadata map[string]string
	// 服务权重
	ServiceWeight int
 
}