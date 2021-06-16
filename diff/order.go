package diff

import "github.com/kong/deck/types"

/*
                                       Root
                                         |
         +----------+----------+---------+------------+---------------+
         |          |          |         |            |               |
         v          v          v         v            v               v
L1    Service    RbacRole  Upstream  Certificate  CACertificate  Consumer ---+
      Package        |         |        |     |      |                |      |
        |            v         v        v     |      v                v      |
L2      |        RBACRole   Target     SNI    +-> Service       Credentials  |
        |        Endpoint                         |  |              (7)      |
        |                                         |  |                       |
        |                                         |  |                       |
L3      +---------------------------> Service <---+  +-> Route               |
        |                             Version        |     |                 |
        |                                 |          |     |                 |
        |                                 |          |     v                 |
L4      +----------> Document   <---------+          +-> Plugins  <----------+
*/

var dependencyOrder = [][]string{
	{
		types.ServicePackage,
		types.RBACRole,
		types.Upstream,
		types.Certificate,
		types.CACertificate,
		types.Consumer,
	},
	{
		types.RBACEndpointPermission,
		types.Target,
		types.SNI,
		types.Service,

		types.KeyAuth, types.HMACAuth, types.JWTAuth,
		types.BasicAuth, types.OAuth2Cred, types.ACLGroup,
		types.MTLSAuth,
	},
	{
		types.ServiceVersion,
		types.Route,
	},
	{
		types.Plugin,
		types.Document,
	},
}

func order() [][]string {
	return deepCopy(dependencyOrder)
}

func reverseOrder() [][]string {
	order := deepCopy(dependencyOrder)
	return reverse(order)
}

func reverse(src [][]string) [][]string {
	src = deepCopy(src)
	i := 0
	j := len(src) - 1
	for i < j {
		temp := src[i]
		src[i] = src[j]
		src[j] = temp
		i++
		j--
	}
	return src
}

func deepCopy(src [][]string) [][]string {
	res := make([][]string, len(src))
	for i := range src {
		res[i] = make([]string, len(src[i]))
		copy(res[i], src[i])
	}
	return res
}
