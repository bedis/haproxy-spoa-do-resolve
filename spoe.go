package main

import (
	spoe "github.com/criteo/haproxy-spoe-go"
)

func doResolve(messages []spoe.Message) (actions []spoe.Action, err error) {
	if debug {
		Debug.Printf("messages: %+v\n", messages)
	}

	for _, msg := range messages {
		var action []spoe.Action

		switch msg.Name {

		case "A":
			var ip string
			var host string
			var ok bool

			if host, ok = msg.Args["host"].(string); ok {
				ip = doResolveA(host)
			}

			action = []spoe.Action{
				spoe.ActionSetVar{
					Name: "addr",
					Scope: spoe.VarScopeTransaction,
					Value: ip,
				},
			}


		case "PTR":
			var addr string
			var name string
			var ok bool

			if addr, ok = msg.Args["addr"].(string); ok {
				name = doResolvePTR(addr)
			}

			action = []spoe.Action{
				spoe.ActionSetVar{
					Name: "hostname",
					Scope: spoe.VarScopeTransaction,
					Value: name,
				},
			}


		default:
			Warning.Printf("Unknonw msg type: %s\n", msg.Name)
			continue
		}

		actions = append(actions, action...)
	}

	if debug {
		Debug.Printf("actions: %+v\n", actions)
	}
	return actions, err
}
