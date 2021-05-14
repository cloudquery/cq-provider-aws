package client

type Error string

func (e Error) Error() string { return string(e) }

const ResourceTypeAssertError = Error("failed to assert resource type")
