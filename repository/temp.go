package repository

/*
Process{
	Def: (*core.Definitions)(0x1400019a000),
	Pool: Pool{
		CustomerSupportIsExecutable: true,
		CustomerIsExecutable: false,
		Collaboration:Builder{Suffix: "0d288062"},
		CustomerSupportID:Builder{Suffix: "ddaf7f1e"},
		CustomerSupportProcess: Builder{Suffix: "2e3ff84f"},
		CustomerID: Builder{Suffix: "199ad245"},
		CustomerProcess: Builder{Suffix: "7df02696"},
	},
	Message: Message{
		CustomerToCustomerSupportMessage: Builder{Suffix: "6acbf202"},
		CustomerSupportToCustomerMessage: Builder{Suffix: "4972f886"},
	},
	CustomerSupport: CustomerSupport{
		CustomerSupportStartEvent: Builder{Suffix: "81c10641"},
		FromCustomerSupportStartEvent: Builder{Suffix: "a7f46176"},
		CheckIncomingClaimTask: Builder{Suffix: "c1926a56"},
		FromCheckIncomingClaimTask: Builder{Suffix: "0850b1d5"},
		DenyWarrantyClaimTask: Builder{Suffix: "5cca5ebf"},
		FromDenyWarrantyClaimTask: Builder{Suffix: "82540b67"},
		CustomerSupportEndEvent: Builder{Suffix: "1730ee62"},
	},
	Customer: Customer{
		CustomerStartEvent: Builder{Suffix: "e5a46e5e"},
		FromCustomerStartEvent: Builder{Suffix: "7206bc52"},
		NoticeOfDefectTask: Builder{Suffix: "f6e8f8f9"},
		FromNoticeOfDefectTask: Builder{Suffix: "cb0a4994"},
		WaitingForAnswerTask: Builder{Suffix: "3a02c5c4"},
		TimerEventDefinitionWaitingForAnswer: Builder{Suffix: "d6549cd4"},
		FromWaitingForAnswerTask: Builder{Suffix: "e7b5a31d"},
		ReceiptWarrantyRefusalTask: Builder{Suffix: "4152c3d5"},
		FromReceiptWarrantyRefusalTask: Builder{Suffix: "6b568430"},
		CustomerEndEvent: Builder{Suffix: "ac518f5e"},
	},
}

Words:

var test = map[int][]string{
	0:  []string{"Pool"},
	1:  []string{"Message"},
	2:  []string{"Customer", "Support"},
	3:  []string{"Customer"},
	4:  []string{"Collaboration"},
	5:  []string{"Customer", "Support", "ID"},
	6:  []string{"Customer", "Support", "Process"},
	7:  []string{"Customer", "ID"},
	8:  []string{"Customer", "Process"},
	9:  []string{"Customer", "To", "Customer", "Support", "Message"},
	10: []string{"Customer", "Support", "To", "Customer", "Message"},
	11: []string{"Customer", "Support", "Start", "Event"},
	12: []string{"From", "Customer", "Support", "Start", "Event"},
	13: []string{"Check", "Incoming", "Claim", "Task"},
	14: []string{"From", "Check", "Incoming", "Claim", "Task"},
	15: []string{"Deny", "Warranty", "Claim", "Task"},
	16: []string{"From", "Deny", "Warranty", "Claim", "Task"},
	17: []string{"Customer", "Support", "End", "Event"},
	18: []string{"Customer", "Start", "Event"},
	19: []string{"From", "Customer", "Start", "Event"},
	20: []string{"Notice", "Of", "Defect", "Task"},
	21: []string{"From", "Notice", "Of", "Defect", "Task"},
	22: []string{"Waiting", "For", "Answer", "Task"},
	23: []string{"Timer", "Event", "Definition", "Waiting", "For", "Answer"},
	24: []string{"From", "Waiting", "For", "Answer", "Task"},
	25: []string{"Receipt", "Warranty", "Refusal", "Task"},
	26: []string{"From", "Receipt", "Warranty", "Refusal", "Task"},
	27: []string{"Customer", "End", "Event"},
	28: []string{"Customer", "Support", "Is", "Executable"},
	29: []string{"Customer", "Is", "Executable"},
}


2023/03/12 09:52:57 factory.builder: start injecting fields
2023/03/12 09:52:57 factory.builder: inject first bool field CustomerSupportIsExecutable once
2023/03/12 09:52:57 factory.builder: next fieldname is CustomerSupportID 1
2023/03/12 09:52:57 factory.builder: previous fieldname was Collaboration
2023/03/12 09:52:57 factory.builder: previous hash value was af9f0585
2023/03/12 09:52:57 factory.builder: next fieldname is CustomerSupportProcess 2
2023/03/12 09:52:57 factory.builder: previous fieldname was CustomerSupportID
2023/03/12 09:52:57 factory.builder: previous hash value was f1475133
2023/03/12 09:52:57 factory.builder: next fieldname is CustomerID 3
2023/03/12 09:52:57 factory.builder: previous fieldname was CustomerSupportProcess
2023/03/12 09:52:57 factory.builder: previous hash value was b6647f6d
2023/03/12 09:52:57 factory.builder: next fieldname is CustomerProcess 4
2023/03/12 09:52:57 factory.builder: previous fieldname was CustomerID
2023/03/12 09:52:57 factory.builder: previous hash value was 51c9592f
2023/03/12 09:52:57 factory.builder: detected field CustomerToCustomerSupportMessage contains anonymous field Message
2023/03/12 09:52:57 factory.builder: next fieldname is CustomerSupportToCustomerMessage 5
2023/03/12 09:52:57 factory.builder: detected field CustomerSupportToCustomerMessage contains anonymous field Message
2023/03/12 09:52:57 factory.builder: previous fieldname was CustomerToCustomerSupportMessage
2023/03/12 09:52:57 factory.builder: previous hash value was be61672a
2023/03/12 09:52:57 factory.builder: detected field CustomerSupportStartEvent contains anonymous field CustomerSupport
2023/03/12 09:52:57 factory.builder: next fieldname is FromCustomerSupportStartEvent 6
2023/03/12 09:52:57 factory.builder: detected field FromCustomerSupportStartEvent contains anonymous field CustomerSupport
2023/03/12 09:52:57 factory.builder: previous fieldname was CustomerSupportStartEvent
2023/03/12 09:52:57 factory.builder: previous hash value was 545b9ad3
2023/03/12 09:52:57 factory.builder: next fieldname is CheckIncomingClaimTask 7
2023/03/12 09:52:57 factory.builder: previous fieldname was FromCustomerSupportStartEvent
2023/03/12 09:52:57 factory.builder: previous hash value was a137afd7
2023/03/12 09:52:57 factory.builder: next fieldname is FromCheckIncomingClaimTask 8
2023/03/12 09:52:57 factory.builder: previous fieldname was CheckIncomingClaimTask
2023/03/12 09:52:57 factory.builder: previous hash value was 409f8399
2023/03/12 09:52:57 factory.builder: next fieldname is DenyWarrantyClaimTask 9
2023/03/12 09:52:57 factory.builder: previous fieldname was FromCheckIncomingClaimTask
2023/03/12 09:52:57 factory.builder: previous hash value was 238fee1a
2023/03/12 09:52:57 factory.builder: next fieldname is FromDenyWarrantyClaimTask 10
2023/03/12 09:52:57 factory.builder: previous fieldname was DenyWarrantyClaimTask
2023/03/12 09:52:57 factory.builder: previous hash value was 26776efb
2023/03/12 09:52:57 factory.builder: next fieldname is CustomerSupportEndEvent 11
2023/03/12 09:52:57 factory.builder: detected field CustomerSupportEndEvent contains anonymous field CustomerSupport
2023/03/12 09:52:57 factory.builder: previous fieldname was FromDenyWarrantyClaimTask
2023/03/12 09:52:57 factory.builder: previous hash value was ce07a934
2023/03/12 09:52:57 factory.builder: detected field CustomerStartEvent contains anonymous field Customer
2023/03/12 09:52:57 factory.builder: next fieldname is FromCustomerStartEvent 12
2023/03/12 09:52:57 factory.builder: detected field FromCustomerStartEvent contains anonymous field Customer
2023/03/12 09:52:57 factory.builder: previous fieldname was CustomerStartEvent
2023/03/12 09:52:57 factory.builder: previous hash value was 75387fdb
2023/03/12 09:52:57 factory.builder: next fieldname is NoticeOfDefectTask 13
2023/03/12 09:52:57 factory.builder: previous fieldname was FromCustomerStartEvent
2023/03/12 09:52:57 factory.builder: previous hash value was f8ccb35b
2023/03/12 09:52:57 factory.builder: next fieldname is FromNoticeOfDefectTask 14
2023/03/12 09:52:57 factory.builder: previous fieldname was NoticeOfDefectTask
2023/03/12 09:52:57 factory.builder: previous hash value was 64a2a3e7
2023/03/12 09:52:57 factory.builder: next fieldname is WaitingForAnswerTask 15
2023/03/12 09:52:57 factory.builder: previous fieldname was FromNoticeOfDefectTask
2023/03/12 09:52:57 factory.builder: previous hash value was cdf67644
2023/03/12 09:52:57 factory.builder: next fieldname is TimerEventDefinitionWaitingForAnswer 16
2023/03/12 09:52:57 factory.builder: previous fieldname was WaitingForAnswerTask
2023/03/12 09:52:57 factory.builder: previous hash value was 1d3bbfab
2023/03/12 09:52:57 factory.builder: next fieldname is FromWaitingForAnswerTask 17
2023/03/12 09:52:57 factory.builder: previous fieldname was TimerEventDefinitionWaitingForAnswer
2023/03/12 09:52:57 factory.builder: previous hash value was b1714bbc
2023/03/12 09:52:57 factory.builder: next fieldname is ReceiptWarrantyRefusalTask 18
2023/03/12 09:52:57 factory.builder: previous fieldname was FromWaitingForAnswerTask
2023/03/12 09:52:57 factory.builder: previous hash value was 6b9b3520
2023/03/12 09:52:57 factory.builder: next fieldname is FromReceiptWarrantyRefusalTask 19
2023/03/12 09:52:57 factory.builder: previous fieldname was ReceiptWarrantyRefusalTask
2023/03/12 09:52:57 factory.builder: previous hash value was 4a748cd9
2023/03/12 09:52:57 factory.builder: next fieldname is CustomerEndEvent 20
2023/03/12 09:52:57 factory.builder: detected field CustomerEndEvent contains anonymous field Customer
2023/03/12 09:52:57 factory.builder: previous fieldname was FromReceiptWarrantyRefusalTask
2023/03/12 09:52:57 factory.builder: previous hash value was e8f4eff4


Process{
	BaseAttributes:impl.BaseAttributes{ID:"Process_f20ac214", Name:""},
	Attributes:attributes.Attributes{
		Documentation:attributes.DOCUMENTATION_SLC(nil),
		ExtensionElements:attributes.EXTENSION_ELEMENTS_SLC(nil)
	},
	ProcessEvents:
		events.ProcessEvents{
			StartEvent:events.START_EVENT_SLC{
				elements.StartEvent{
					BaseAttributes:impl.BaseAttributes{ID:"Event_d316c6c0", Name:""},
					CoreAttributes:camunda.CoreAttributes{CamundaAsyncBefore:false, CamundaAsyncAfter:false, CamundaJobPriority:0},
					Attributes:attributes.Attributes{
						Documentation:attributes.DOCUMENTATION_SLC(nil),
						ExtensionElements:attributes.EXTENSION_ELEMENTS_SLC(nil)
					},
					StartEvent:definitions.StartEvent{
						ConditionalEventDef:definitions.CONDITIONAL_EVENT_DEF_SLC(nil),
						MessageEventDefinition:definitions.MESSAGE_EVENT_DEF_SLC(nil),
						TimerEventDef:definitions.TIMER_EVENT_DEF_SLC(nil)
					},
					IsInterrupting:false,
					CamundaFormKey:"",
					CamundaFormRef:"",
					CamundaFormRefBind:"",
					CamundaFormRefVersion:"",
					CamundaInitiator:"",
					Outgoing:[]marker.Outgoing{marker.Outgoing{Flow:"Flow_1f2c5dbb"}}
				}
			},
			BoundaryEvent:events.BOUNDARY_EVENT_SLC(nil),
			EndEvent:events.END_EVENT_SLC{
				elements.EndEvent{
					BaseAttributes:impl.BaseAttributes{ID:"", Name:""},
					CoreAttributes:camunda.CoreAttributes{CamundaAsyncBefore:false, CamundaAsyncAfter:false, CamundaJobPriority:0},
					Attributes:attributes.Attributes{Documentation:attributes.DOCUMENTATION_SLC(nil), ExtensionElements:attributes.EXTENSION_ELEMENTS_SLC(nil)},
					EndEvent:definitions.EndEvent{
						CompensateEventDefinition:definitions.COMPENSATE_EVENT_DEF_SLC(nil),
						EscalationEventDefinition:definitions.ESCALATION_EVENT_DEF_SLC(nil),
						MessageEventDefinition:definitions.MESSAGE_EVENT_DEF_SLC(nil),
						ErrorEventDefinition:definitions.ERROR_EVENT_DEF_SLC(nil),
						SignalEventDefinition:definitions.SIGNAL_EVENT_DEF_SLC(nil),
						TerminateEventDefinition:definitions.TERMINATE_EVENT_DEF_SLC(nil)
					},
					Incoming:[]marker.Incoming(nil)
				}
			},
			IntermediateCatchEvent:events.INTERMEDIATE_CATCH_EVENT_SLC(nil),
			IntermediateThrowEvent:events.INTERMEDIATE_THROW_EVENT_SLC(nil)
		},
		Tasks:tasks.Tasks{
			BusinessRuleTask:tasks.BUSINESS_RULE_TASK_SLC(nil),
			Task:tasks.TASK_SLC{
				tasks.Task{
					BaseAttributes:impl.BaseAttributes{ID:"", Name:""},
					Attributes:attributes.Attributes{Documentation:attributes.DOCUMENTATION_SLC(nil), ExtensionElements:attributes.EXTENSION_ELEMENTS_SLC(nil)},
					CoreAttributes:camunda.CoreAttributes{CamundaAsyncBefore:false, CamundaAsyncAfter:false, CamundaJobPriority:0},
					IncomingOutgoing:marker.IncomingOutgoing{Incoming:[]marker.Incoming(nil), Outgoing:[]marker.Outgoing(nil)},
					Property:[]attributes.Property(nil),
					DataInputAssociation:[]flow.DataInputAssociation(nil),
				},
				tasks.Task{
					BaseAttributes:impl.BaseAttributes{ID:"", Name:""},
					Attributes:attributes.Attributes{Documentation:attributes.DOCUMENTATION_SLC(nil), ExtensionElements:attributes.EXTENSION_ELEMENTS_SLC(nil)},
					CoreAttributes:camunda.CoreAttributes{CamundaAsyncBefore:false, CamundaAsyncAfter:false, CamundaJobPriority:0},
					IncomingOutgoing:marker.IncomingOutgoing{Incoming:[]marker.Incoming(nil), Outgoing:[]marker.Outgoing(nil)},
					Property:[]attributes.Property(nil),
					DataInputAssociation:[]flow.DataInputAssociation(nil)
				}
			},
			UserTask:tasks.USER_TASK_SLC(nil),
			ManualTask:tasks.MANUAL_TASK_SLC(nil),
			ReceiveTask:tasks.RECEIVE_TASK_SLC(nil),
			ScriptTask:tasks.SCRIPT_TASK_SLC(nil),
			SendTask:tasks.SEND_TASK_SLC(nil),
			ServiceTask:tasks.SERVICE_TASK_SLC(nil)
		},
		Subprocesses:subprocesses.Subprocesses{
			CallActivity:subprocesses.CALL_ACTIVITY_SLC(nil),
			SubProcess:subprocesses.SUBPROCESS_SLC(nil),
			Transaction:subprocesses.TRANSACTION_SLC(nil),
			AdHocSubProcess:subprocesses.ADHOC_SUBPROCESS_SLC(nil)
		},
		Gateways:gateways.Gateways{
			ExclusiveGateway:gateways.EXCLUSIVE_GATEWAY_SLC(nil),
			InclusiveGateway:gateways.INCLUSIVE_GATEWAY_SLC(nil),
			ParallelGateway:gateways.PARALLEL_GATEWAY_SLC(nil),
			ComplexGateway:gateways.COMPLEX_GATEWAY_SLC(nil),
			EventBasedGateway:gateways.EVENT_BASED_GATEWAYS_SLC(nil)
		},
		IsExecutable:true,
		CamundaVersionTag:"",
		CamundaJobPriority:0,
		CamundaTaskPriority:0,
		CamundaCandidateStarterGroups:"",
		CamundaCandidateStarterUsers:"",
		CamundaHistoryTimeToLive:"",
		LaneSet:[]pool.LaneSet(nil),
		Association:[]flow.Association(nil),
		SequenceFlow:[]flow.SequenceFlow{
			flow.SequenceFlow{
				BaseAttributes:impl.BaseAttributes{ID:"", Name:""},
				Attributes:attributes.Attributes{Documentation:attributes.DOCUMENTATION_SLC(nil), ExtensionElements:attributes.EXTENSION_ELEMENTS_SLC(nil)},
				SourceTargetRef:flow.SourceTargetRef{SourceRef:"", TargetRef:""},
				ConditionExpression:[]conditional.ConditionExpression(nil)
			},
			flow.SequenceFlow{
				BaseAttributes:impl.BaseAttributes{ID:"", Name:""},
				Attributes:attributes.Attributes{Documentation:attributes.DOCUMENTATION_SLC(nil), ExtensionElements:attributes.EXTENSION_ELEMENTS_SLC(nil)},
				SourceTargetRef:flow.SourceTargetRef{SourceRef:"", TargetRef:""},
				ConditionExpression:[]conditional.ConditionExpression(nil)
			},
			flow.SequenceFlow{
				BaseAttributes:impl.BaseAttributes{ID:"", Name:""},
				Attributes:attributes.Attributes{Documentation:attributes.DOCUMENTATION_SLC(nil), ExtensionElements:attributes.EXTENSION_ELEMENTS_SLC(nil)},
				SourceTargetRef:flow.SourceTargetRef{SourceRef:"", TargetRef:""},
				ConditionExpression:[]conditional.ConditionExpression(nil)
			}
		},
		Group:[]marker.Group(nil),
		DataObject:[]data.DataObject(nil)
}

*/
