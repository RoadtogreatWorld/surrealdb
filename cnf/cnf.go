// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cnf

import (
	"net"
	"time"
)

var Settings *Options

type Kind int

func (k Kind) String() string {
	switch k {
	default:
		return "NO"
	case AuthKV:
		return "KV"
	case AuthNS:
		return "NS"
	case AuthDB:
		return "DB"
	case AuthSC:
		return "SC"
	}
}

const (
	// Root access
	AuthKV Kind = iota
	// Namespace access
	AuthNS
	// Database access
	AuthDB
	// Scoped user access
	AuthSC
	// No access
	AuthNO
)

type Auth struct {
	Kind     Kind
	Data     map[string]interface{}
	Possible struct {
		NS string
		DB string
	}
	Selected struct {
		NS string
		DB string
	}
}

// Options defines global configuration options
type Options struct {
	DB struct {
		Key  []byte        // Data encryption key
		Code string        // Data encryption key string
		Path string        // Path to store the data file
		Host string        // Surreal host to connect to
		Port string        // Surreal port to connect to
		Base string        // Base key to use in KV stores
		Time time.Duration // Timeframe for syncing data
		Cert struct {
			CA  string
			Crt string
			Key string
			SSL bool
		}
	}

	Port struct {
		Web int // Web port as an number
		Tcp int // Tcp port as an number
	}

	Conn struct {
		Web string // Web port as a string
		Tcp string // Tcp port as a string
	}

	Cert struct {
		Crt string // File location of server crt
		Key string // File location of server key
	}

	Auth struct {
		Auth string       // Master authentication user:pass
		User string       // Master authentication username
		Pass string       // Master authentication password
		Addr []string     // Allowed ip ranges for authentication
		Nets []*net.IPNet // Allowed cidr ranges for authentication
	}

	Node struct {
		Host string   // Node hostname
		Name string   // Name of this node
		UUID string   // UUID of this node
		Join []string // Slice of cluster peers to join
	}

	Logging struct {
		Level  string // Stores the configured logging level
		Output string // Stores the configured logging output
		Format string // Stores the configured logging format
		Google struct {
			Name        string // Stores the GCE logging name
			Project     string // Stores the GCE logging project
			Credentials string // Store the path to the credentials file
		}
		Syslog struct {
			Tag      string // Stores the syslog tag name
			Host     string // Stores the syslog remote host:port
			Protocol string // Stores the syslog protocol to use
			Priority string // Stores the syslog logging priority
		}
	}
}
