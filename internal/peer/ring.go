package peer

type Ring struct {
	self string
	peers *consistent.Consistent
}

func NewRing(self string, peerList []string) *Ring {
	r := &Ring {
		self: self,
		peers: consistent.New()
	}

	for _, p := range peerList {
		if p != self {
			r.peers.Add(p)
		}
	}
	return r
}
