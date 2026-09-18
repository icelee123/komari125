package agent

import (
	"testing"

	v2 "github.com/icelee123/komari125/protocol/v2"
)

func TestV2EventWhitelist(t *testing.T) {
	const uuid = "test-v2-event-whitelist"

	if event := EnqueueV2Event(uuid, "agent.exec", map[string]any{"command": "id"}); event.ID != "" {
		t.Fatal("expected remote command event to be rejected")
	}
	if event := EnqueueV2Event(uuid, v2.MethodAgentPing, v2.PingParams{}); event.ID == "" {
		t.Fatal("expected ping event to be accepted")
	}

	v2EventMu.Lock()
	delete(v2EventQueues, uuid)
	v2EventMu.Unlock()
}
