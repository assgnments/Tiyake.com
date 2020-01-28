package permission

import (
	"strings"
)

type permission struct {
	roles   []string
	methods []string
}

type authority map[string]permission

var authorities = authority{
	"/": permission{
		roles:   []string{"USER", "ADMIN"},
		methods: []string{"POST", "GET"},
	},
	"/question": permission{
		roles:   []string{"USER", "ADMIN"},
		methods: []string{"POST", "GET"},
	},
	"/question/new": permission{
		roles:   []string{"USER"},
		methods: []string{"POST", "GET"},
	},
	"/question/upvote": permission{
		roles:   []string{"USER"},
		methods: []string{"POST", "GET"},
	},
	"/question/search": permission{
		roles:   []string{"USER"},
		methods: []string{"POST", "GET"},
	},
	"/api/questions": permission{
		roles:   []string{"USER"},
		methods: []string{"POST", "GET"},
	},
	"/logout": permission{
		roles:   []string{"USER", "ADMIN"},
		methods: []string{"POST"},
	},
	"/admin": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
	"/admin/users": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
	"/admin/users/update": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
	"/admin/users/new": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
	"/admin/users/delete": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
	"/admin/categores": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
	"/admin/categories/update": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
	"/admin/categories/delete": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
	"/admin/categories/new": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
	"/admin/questions": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
	"/admin/questions/delete": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
	"/admin/answers": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
	"/admin/answers/delete": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
}

// HasPermission checks if a given role has permission to access a given route for a given method
func HasPermission(role string, path string, method string) bool {
	if strings.HasPrefix(path, "/admin") {
		path = "/admin"

	}
	perm := authorities[path]
	checkedRole := checkRole(role, perm.roles)
	checkedMethod := checkMethod(method, perm.methods)
	if checkedRole && checkedMethod {
		return true
	}
	return false
}

func checkRole(role string, roles []string) bool {
	for _, r := range roles {
		if strings.ToUpper(r) == strings.ToUpper(role) {
			return true
		}
	}
	return false
}

func checkMethod(method string, methods []string) bool {
	for _, m := range methods {
		if strings.ToUpper(m) == strings.ToUpper(method) {
			return true
		}
	}
	return false
}
