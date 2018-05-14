package routes

import (
	
)

// RegisterInstall ...
// when install yunus platform, we split the step into four parts
// a. config and check healthz
// b. preflight
// c. postflight
// d. complete check
// we not only provide to install a new cluster, but a new node can be
// join to a current cluster.
