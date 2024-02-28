# Test

## Collaborative Process

```json
{
    Reflect{
    Analyze:Analyze{
        Anonym:map[int]string{
            0:"Pool",
            1:"Message",
            2:"CustomerSupport",
            3:"Customer"
        },
        Config:map[int]string{
            0:"CustomerSupportIsExecutable",
            1:"CustomerIsExecutable"
        },
        Builder:map[int]string{
            0:"Collaboration",
            1:"CustomerSupportID",
            2:"CustomerSupportProcess",
            3:"CustomerID", 
            4:"CustomerProcess", 
            5:"CustomerToCustomerSupportMessage", 
            6:"CustomerSupportToCustomerMessage", 
            7:"CustomerSupportStartEvent", 
            8:"FromCustomerSupportStartEvent", 
            9:"CheckIncomingClaimTask", 
            10:"FromCheckIncomingClaimTask", 
            11:"DenyWarrantyClaimTask", 
            12:"FromDenyWarrantyClaimTask", 
            13:"CustomerSupportEndEvent", 
            14:"CustomerStartEvent", 
            15:"FromCustomerStartEvent", 
            16:"NoticeOfDefectTask", 
            17:"FromNoticeOfDefectTask", 
            18:"WaitingForAnswerTask", 
            19:"TimerEventDefinitionWaitingForAnswer", 
            20:"FromWaitingForAnswerTask", 
            21:"ReceiptWarrantyRefusalTask", 
            22:"FromReceiptWarrantyRefusalTask",
            23:"CustomerEndEvent"
        }, 
        Words:map[int][]string{
            0:[]string{"Pool"},
            1:[]string{"Message"},
            2:[]string{"Customer", "Support"},
            3:[]string{"Customer"},
            4:[]string{"Collaboration"}, 
            5:[]string{"Customer", "Support", "ID"}, 
            6:[]string{"Customer", "Support", "Process"}, 
            7:[]string{"Customer", "ID"}, 
            8:[]string{"Customer", "Process"}, 
            9:[]string{"Customer", "To", "Customer", "Support", "Message"}, 
            10:[]string{"Customer", "Support", "To", "Customer", "Message"}, 
            11:[]string{"Customer", "Support", "Start", "Event"}, 
            12:[]string{"From", "Customer", "Support", "Start", "Event"}, 
            13:[]string{"Check", "Incoming", "Claim", "Task"}, 
            14:[]string{"From", "Check", "Incoming", "Claim", "Task"}, 
            15:[]string{"Deny", "Warranty", "Claim", "Task"}, 
            16:[]string{"From", "Deny", "Warranty", "Claim", "Task"}, 
            17:[]string{"Customer", "Support", "End", "Event"}, 
            18:[]string{"Customer", "Start", "Event"}, 
            19:[]string{"From", "Customer", "Start", "Event"}, 
            20:[]string{"Notice", "Of", "Defect", "Task"}, 
            21:[]string{"From", "Notice", "Of", "Defect", "Task"}, 
            22:[]string{"Waiting", "For", "Answer", "Task"}, 
            23:[]string{"Timer", "Event", "Definition", "Waiting", "For", "Answer"}, 
            24:[]string{"From", "Waiting", "For", "Answer", "Task"}, 
            25:[]string{"Receipt", "Warranty", "Refusal", "Task"}, 
            26:[]string{"From", "Receipt", "Warranty", "Refusal", "Task"}, 
            27:[]string{"Customer", "End", "Event"}, 
            28:[]string{"Customer", "Support", "Is", "Executable"}, 
            29:[]string{"Customer", "Is", "Executable"}
        }
    }, 
    IF:Process{
        Def:core.DefinitionsRepository(nil), 
        Pool:collaborative_process.Pool{
            CustomerSupportIsExecutable:false, 
            CustomerIsExecutable:false, 
            Collaboration:factory.Builder{
                Reflect:factory.Reflect{
                    Analyze:factory.Analyze{
                        Anonym:map[int]string(nil), 
                        Config:map[int]string(nil), 
                        Builder:map[int]string(nil), 
                        Words:map[int][]string(nil)
                    }, 
                    IF:interface {}(nil), 
                    Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, 
                    Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}
                }, 
                Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, 
                Suffix:""
            }, 
            CustomerSupportID:factory.Builder{
                Reflect:factory.Reflect{
                        Analyze:factory.Analyze{Anonym:map[int]string(nil), 
                        Config:map[int]string(nil), 
                        Builder:map[int]string(nil), 
                        Words:map[int][]string(nil)
                }, 
                    IF:interface {}(nil), 
                    Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, 
                    Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}
                }, 
                Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, 
                Suffix:""
            }, 
            CustomerSupportProcess:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}, 
            CustomerID:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}, 
            CustomerProcess:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}
        }, 
        Message:collaborative_process.Message{
            CustomerToCustomerSupportMessage:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}, 
            CustomerSupportToCustomerMessage:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}
        }, 
        CustomerSupport:collaborative_process.CustomerSupport{
            CustomerSupportStartEvent:factory.Builder{
                Reflect:factory.Reflect{
                    Analyze:factory.Analyze{
                        Anonym:map[int]string(nil), 
                        Config:map[int]string(nil), 
                        Builder:map[int]string(nil), 
                        Words:map[int][]string(nil)
                    }, 
                    IF:interface {}(nil), 
                    Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, 
                    Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}
                }, 
                Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, 
                Suffix:""
            }, 
            FromCustomerSupportStartEvent:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}, 
            CheckIncomingClaimTask:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}, 
            FromCheckIncomingClaimTask:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}, 
            DenyWarrantyClaimTask:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}, 
            FromDenyWarrantyClaimTask:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}, 
            CustomerSupportEndEvent:factory.Builder{
                Reflect:factory.Reflect{
                    Analyze:factory.Analyze{
                        Anonym:map[int]string(nil), 
                        Config:map[int]string(nil), 
                        Builder:map[int]string(nil), 
                        Words:map[int][]string(nil)
                    }, 
                    IF:interface {}(nil), 
                    Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, 
                    Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}
                }, 
                Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, 
                Suffix:""
            }
        },
        Customer:collaborative_process.Customer{
            CustomerStartEvent:factory.Builder{
                Reflect:factory.Reflect{
                    Analyze:factory.Analyze{
                        Anonym:map[int]string(nil), 
                        Config:map[int]string(nil), 
                        Builder:map[int]string(nil), 
                        Words:map[int][]string(nil)
                    }, 
                    IF:interface {}(nil), 
                    Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, 
                    Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}
                }, 
                Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, 
                Suffix:""
            },
            FromCustomerStartEvent:factory.Builder{
                Reflect:factory.Reflect{
                    Analyze:factory.Analyze{
                        Anonym:map[int]string(nil), 
                        Config:map[int]string(nil), 
                        Builder:map[int]string(nil), 
                        Words:map[int][]string(nil)
                    }, 
                    IF:interface {}(nil), 
                    Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, 
                    Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}
                }, 
                Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, 
                Suffix:""
            }, 
            NoticeOfDefectTask:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}, 
            FromNoticeOfDefectTask:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}, 
            WaitingForAnswerTask:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}, 
            TimerEventDefinitionWaitingForAnswer:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}, 
            FromWaitingForAnswerTask:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}, 
            ReceiptWarrantyRefusalTask:factory.Builder{Reflect:factory.Reflect{Analyze:factory.Analyze{Anonym:map[int]string(nil), Config:map[int]string(nil), Builder:map[int]string(nil), Words:map[int][]string(nil)}, IF:interface {}(nil), Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}}, Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, Suffix:""}, 
            FromReceiptWarrantyRefusalTask:factory.Builder{
                Reflect:factory.Reflect{
                    Analyze:factory.Analyze{
                        Anonym:map[int]string(nil), 
                        Config:map[int]string(nil), 
                        Builder:map[int]string(nil), 
                        Words:map[int][]string(nil)
                    }, 
                    IF:interface {}(nil), 
                    Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, 
                    Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}
                }, 
                Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, 
                Suffix:""
            }, 
            CustomerEndEvent:factory.Builder{
                Reflect:factory.Reflect{
                    Analyze:factory.Analyze{
                        Anonym:map[int]string(nil), 
                        Config:map[int]string(nil), 
                        Builder:map[int]string(nil), 
                        Words:map[int][]string(nil)
                    }, 
                    IF:interface {}(nil), 
                    Temporary:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}, 
                    Element:reflect.Value{typ:(*reflect.rtype)(nil), ptr:(unsafe.Pointer)(nil), flag:0x0}
                },
                Settings:factory.Settings{NumProc:0, NumPart:0, NumMsg:0, NumStartEvent:0, NumEndEvent:0, NumTask:0, NumFlow:0, NumShape:0, NumEdge:0}, 
                Suffix:""
            }
        }
    }, 
    Temporary:reflect.Value{typ:(*reflect.rtype)(0x100b3e960), ptr:(unsafe.Pointer)(0x14000182600), flag:0x199}, 
    Element:reflect.Value{typ:(*reflect.rtype)(0x100af5180), ptr:(unsafe.Pointer)(0x140001102c0), flag:0x194}
    }
}
```

