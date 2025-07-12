package payment

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	NotificationActionImp = NotificationAction
}
func NotificationAction(
	req *NotificationActionReqDto,
	q fireback.QueryDSL) (string,
	*fireback.IError,
) {
	// Implement the logic here.
	return "", nil
}
