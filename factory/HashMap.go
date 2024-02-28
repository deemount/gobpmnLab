package factory

type AnonymMap struct {
	Field int
	Value []string
}

// hashTable ...
// Note:
// @self is a pseudo and means "this hash is set to this element"
// Ruleset for a straight process without any decision:
// A Collaboration has one hash value: self
// A Participant has two hash values: id, process
// A Pool has one hash value: id
// A StartEvent has two hash values: self, next (flow: from)
// A Flow has three hash values: self, previous, next (task ...)
// A Task, Event has three hash values: self, previous (flow: from), next (flow. from)
// A EndEvent has two hash values: self, previous (flow: from)

/*

map[
	Customer:map[
		0:[CustomerStartEvent]
		1:[FromCustomerStartEvent]
		2:[NoticeOfDefectTask]
		3:[FromNoticeOfDefectTask]
		4:[WaitingForAnswerTask]
		5:[TimerEventDefinitionWaitingForAnswer]
		6:[FromWaitingForAnswerTask]
		7:[ReceiptWarrantyRefusalTask]
		8:[FromReceiptWarrantyRefusalTask]
		9:[CustomerEndEvent]
	]
	CustomerSupport:map[
		0:[CustomerSupportStartEvent]
		1:[FromCustomerSupportStartEvent]
		2:[CheckIncomingClaimTask]
		3:[FromCheckIncomingClaimTask]
		4:[DenyWarrantyClaimTask]
		5:[FromDenyWarrantyClaimTask]
		6:[CustomerSupportEndEvent]
	]
	Message:map[
		0:[CustomerToCustomerSupportMessage]
		1:[CustomerSupportToCustomerMessage]
	]
	Pool:map[
		0:[CustomerSupportIsExecutable]
		1:[CustomerIsExecutable]
		2:[Collaboration]
		3:[CustomerSupportID]
		4:[CustomerSupportProcess]
		5:[CustomerID]
		6:[CustomerProcess]
	]
]

*/
