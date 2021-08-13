// Package ytoken provides a token authentication scheme based on yaml style.
//
// YToken Format:
// 	  Protocol:
//   	- algo: xx			// 签名算法
//   	- type: token		// 协议类型
// 	  Constraint:
//		- signer: xx		// 签发者
//		- expiry: xx		// 过期时间
//		- serial: xx		// 序列号
//		- beneficiary: xx	// 受惠者
//
//	Token Format:
//		${base64-yaml-text}-${base64-hmac}
//
package ytoken
