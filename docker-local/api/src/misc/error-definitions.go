// Package misc defines miscellaneous functions
package misc

import (
	"net/http"
)

// 事前定義済みエラー一覧
var (
	//COMMON系 ERROR ------------------------------------------------------------------
	// 現在サービスは利用できません。しばらく時間を置いてから、もう一度お試しください。
	ServiceUnavailable = newError(http.StatusServiceUnavailable, "E100000", "error.unavailable")
	// 予期しないエラーが発生しました。しばらく時間を置いてから、もう一度お試しください。
	InternalServerError = newError(http.StatusInternalServerError, "E100001", "error.internal")
	// 見つかりません
	NotFound = newError(http.StatusNotFound, "E404", "error.not-found")
	// すでに登録されています
	Duplicated = newError(http.StatusConflict, "E409", "error.duplicated")
	// 認証が必要です。
	Unauthorized = newError(http.StatusUnauthorized, "E100002", "error.unauthorized")
	// JWTが不正です。
	InvalidJSONWebToken = newError(http.StatusUnauthorized, "E100003", "error.invalid-jwt")
	// 不正な入力値で処理を続行できません。
	InvalidParameters = newError(http.StatusBadRequest, "E100004", "error.invalid-parameters")
	// トークンが不正です。
	InvalidToken = newError(http.StatusBadRequest, "E100005", "error.invalid-token")
	// アクセス権限がありません
	Forbidden = newError(http.StatusForbidden, "E100006", "error.forbidden")
	// 有効期限切れです。
	InvalidExpired = newError(http.StatusBadRequest, "E100006", "error.invalid-expired")
	// Eメールアドレス、またはパスワードが不正です。
	InvalidCredentials = newError(http.StatusUnauthorized, "E102101", "error.invalid-credentials")
	// 認証データの作成に失敗しました。
	CreateAuthdataError = newError(http.StatusInternalServerError, "E102102", "error.create-authdata")
	// ユーザー
	UserNotFoundError = newError(http.StatusNotFound, "E102001", "error.user-not-found")
	// パスワードが一致しないまたは形式が違います。
	UserChangePasswordError = newError(http.StatusBadRequest, "E102103", "error.user-change-password-error")
	// パスワードが一致しないまたは形式が違います。
	UserFormatPasswordError = newError(http.StatusBadRequest, "E102106", "error.user-format-password-error")
	// そのメールアドレスは既に使用されています
	UserExistEmailError = newError(http.StatusBadRequest, "E102104", "error.user-exist-email-error")
	// そのニックネームは既に使用されています
	UserExistNicknameError = newError(http.StatusBadRequest, "E102105", "error.user-exist-nickname-error")
	// 既に登録済みの電話番号です
	UserExistPhoneError = newError(http.StatusBadRequest, "E102107", "error.user-exist-nickname-error")
	// 支援グループが無効
	SupportGroupNotFoundError = newError(http.StatusNotFound, "E102010", "error.support-group-not-found")
	//マーチャントが無効
	MerchantNotFoundError = newError(http.StatusNotFound, "E102002", "error.merchant-not-found")
	//マーチャントが無効
	MerchantOwnerNotFoundError = newError(http.StatusNotFound, "E102014", "error.merchant-owner-not-found")
	//既に事業者オーナーとして登録されているユーザーです
	MerchantAlreadyUserError = newError(http.StatusBadRequest, "E102011", "error.merchant-already-user-error")
	//プロジェクトが無効
	ProjectNotFoundError = newError(http.StatusNotFound, "E102003", "error.project-not-found")
	//プロジェクトプランが無効
	ProjectPlanNotFoundError = newError(http.StatusNotFound, "E102004", "error.project-plan-not-found")
	//プロジェクト画像が無効
	ProjectImageNotFoundError = newError(http.StatusNotFound, "E102012", "error.project-image-not-found")
	//プロジェクト画像が無効
	ProjectImageUploadLimitError = newError(http.StatusBadRequest, "E102013", "error.project-image-upload-limit")
	//注文情報が無効
	OrderNotFoundError = newError(http.StatusNotFound, "E102005", "error.order-not-found")
	//注文寄付情報が無効
	OrderDonationNotFoundError = newError(http.StatusNotFound, "E102006", "error.order-donation-not-found")
	//注文クーポン情報が無効
	OrderCouponNotFoundError = newError(http.StatusNotFound, "E102007", "error.order-coupon-not-found")
	//注文クーポン情報の有効期限切れ
	ExpiredOrderCouponError = newError(http.StatusBadRequest, "E102008", "error.expired-order-coupon")
	//注文クーポンが使用済み
	UsedOrderCouponError = newError(http.StatusBadRequest, "E102008", "error.used-order-coupon")
	// ユーザー住所が無効
	UserAddressNotFoundError = newError(http.StatusNotFound, "E102009", "error.user-address-not-found")
	// 決済処理が正常に行われませんでした。
	ExecuteSettlementError = newError(http.StatusInternalServerError, "E104001", "error.execute-settlement")
	// 検証コードの有効期限切れです。最初からメールアドレス変更をお願いします。
	MailInvalidExpired = newError(http.StatusBadRequest, "E103001", "error.mail-invalid-expired")
	// 検証回数の上限に達しました。最初からメールアドレス変更をお願いします。
	MailLimitExceeded = newError(http.StatusBadRequest, "E103002", "error.mail-limit-exceeded")
	// クレジットカードが定期支援で使用中です
	CardUsedForSubscription = newError(http.StatusBadRequest, "E104002", "error.card-is-used-for-subscription")
	//店舗コードが違うまたは入力された店舗コードではクーポンが使用できません
	StoreOrderCouponError = newError(http.StatusBadRequest, "E102015", "error.store-order-coupon")
	// 在庫が存在しません
	StockEmptyError = newError(http.StatusBadRequest, "E102015", "error.stock-empty")
)
