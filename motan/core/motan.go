package core

type Serialization interface {
	GetSerialNum() int
	Serialize(v interface{}) ([]byte, error)
	DeSerialize(b []byte, v interface{}) (interface{}, error)
	SerializeMulti(v []interface{}) ([]byte, error)
	DeSerializeMulti(b []byte, v []interface{}) ([]interface{}, error)
}

type ExtensionFactory interface {
	GetHa(url *URL) HaStrategy
}

type Name interface {
	GetName() string
}

type Identity interface {
	GetIdentity() string
}

type Destroyable interface {
	Destroy()
}

type Server interface {
	WithURL
	Name
	Destroyable
}

type Exporter interface {
	Export(server Server, factory ExtensionFactory)
}

type WithURL interface {
	GetURL() *URL
	SetURL(url *URL)
}

type Attachment interface {
	GetAttachment()
}

type HaStrategy interface {
}

type Initializable interface {
	Initialize()
}

type NotifyListener interface {
	Identity
	Notify(registryURL *URL, url []*URL)
}

type DiscoverService interface {
	Subscribe(url *URL, listener NotifyListener)
	Unsubscribe(url *URL, listener NotifyListener)
	Discover(url *URL) []*URL
}

type Registry interface {
	Name
	WithURL
}

func Initialize(s interface{}) {
	if init, ok := s.(Initializable); ok {
		init.Initialize()
	}
}
