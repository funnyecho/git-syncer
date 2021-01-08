package gitrepo

func (r *repo) GetHead() (string, error) {
	var head string

	if symbolicHead := r.gitter.GetSymbolicHead(); symbolicHead != "" {
		head = symbolicHead
	} else if headRev := r.gitter.GetHead(); headRev != "" {
		head = headRev
	} else if localSHA1, localSHA1Err := r.gitter.GetHeadSHA1(); localSHA1Err == nil {
		head = localSHA1
	}

	return head, nil
}

func (r *repo) GetHeadSHA1() (string, error) {
	return r.gitter.GetHeadSHA1()
}

func (r *repo) PushHead(head string) (string, error) {
	prevHead, prevHeadErr := r.GetHead()

	if prevHeadErr != nil {
		return "", prevHeadErr
	}

	checkoutErr := r.gitter.Checkout(head)
	if checkoutErr != nil {
		return "", checkoutErr
	}

	return prevHead, nil
}

func (r *repo) PopHead(head string) error {
	return r.gitter.Checkout(head)
}
