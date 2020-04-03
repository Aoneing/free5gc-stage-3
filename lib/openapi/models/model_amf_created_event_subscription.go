/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AmfCreatedEventSubscription struct {
	Subscription      *AmfEventSubscription `json:"subscription" bson:"subscription" `
	SubscriptionId    string                `json:"subscriptionId" bson:"subscriptionId" `
	ReportList        []AmfEventReport      `json:"reportList,omitempty" bson:"reportList" `
	SupportedFeatures string                `json:"supportedFeatures,omitempty" bson:"supportedFeatures" `
}