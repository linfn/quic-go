package ackhandler

import "github.com/lucas-clemente/quic-go/internal/wire"

// IsFrameAckEliciting returns true if the frame is ack-eliciting.
func IsFrameAckEliciting(f wire.Frame) bool {
	_, ok := f.(*wire.AckFrame)
	return !ok
}

// HasAckElicitingFrames returns true if at least one frame is ack-eliciting.
func HasAckElicitingFrames(fs []wire.Frame) bool {
	for _, f := range fs {
		if IsFrameAckEliciting(f) {
			return true
		}
	}
	return false
}
