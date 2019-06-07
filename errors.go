package mcache

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrNotFound = status.Error(codes.NotFound, "cache entry not found")
