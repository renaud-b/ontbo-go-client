# Ontbo Go Client

This repository provides a **Go client** for the [Ontbo API](https://www.ontbo.com).  
It implements all available endpoints from the official OpenAPI specification (profiles, scenes, facts, context, update).  

---

## 🚀 Installation

Clone the repository and import the `ontbo` package in your Go project:

```bash
go get github.com/renaud-b/ontbo-go-client
```

---

## 🔑 Authentication

You need an API token (Bearer) from the [Ontbo developer hub](https://api.ontbo.com/).  

Set it in your environment:

```bash
export ONTBO_TOKEN=your_token_here
```

---

## 📦 Usage Examples

### 1. Initialize the client

```go
import "github.com/yourusername/ontbo"

client := ontbo.NewClient("https://api.ontbo.com/api/tests", os.Getenv("ONTBO_TOKEN"))
```

---

### 2. Profiles

```go
// List profiles
profiles, _ := client.ListProfiles()
fmt.Println(profiles)

// Create a new profile
profile, _ := client.CreateProfile("test_profile")
fmt.Println("New profile:", profile.ID)

// Delete a profile
err := client.DeleteProfile(profile.ID)
```

---

### 3. Facts

```go
// Add a fact
fact, _ := client.AddFact(profile.ID, "I love programming", "manual_test")
fmt.Println("Added fact:", fact.Data)

// List facts
facts, _ := client.ListFacts(profile.ID, nil, 0, 10)
for _, f := range facts {
    fmt.Println("Fact:", f.Data)
}

// Query facts
res, _ := client.QueryFacts(profile.ID, "What do I love?", "VECTOR_SEARCH")
fmt.Println("Query result:", res.Answer)
```

---

### 4. Scenes

```go
// Create a scene
scene, _ := client.CreateScene(profile.ID, "chat1")

// Add text to a scene
messages := []ontbo.SceneMessage{
    {Role: "assistant", Content: "Hello!", Timestamp: float64(time.Now().UnixNano())/1e9},
    {Role: "user", Content: "I am 30 years old.", Timestamp: float64(time.Now().UnixNano())/1e9},
}
updated, _ := client.AddTextToScene(profile.ID, scene.ID, messages, true, true)
fmt.Println("Updated scene:", updated)

// Query scenes
res, _ := client.QueryScenes(profile.ID, "How old am I?")
fmt.Println("Scene query:", res.Answer)
```

---

### 5. Context

```go
// Build context from a query
ctx, _ := client.BuildContext(profile.ID, "Tell me about my animals")
fmt.Println("Context:", ctx.Context)
```

---

### 6. Updates

```go
// Run update
_, _ = client.RunProfileUpdate(profile.ID)

// Check status
status, _ := client.GetProfileUpdateStatus(profile.ID)
fmt.Println("Status:", status.Status, "Progress:", status.Progress)

// Stop update
_, _ = client.StopProfileUpdate(profile.ID)
```

---

## 🧩 Features

- **Profiles**: create, list, delete, update  
- **Facts**: add, list, query, delete  
- **Scenes**: create, add text, query, delete  
- **Context**: build contextual responses  
- **Updates**: run/stop consolidation and check status  

---

## ⚠️ Notes

- The API is evolving, some endpoints may change.  
- Not all response fields are fully documented yet, so check responses when integrating.  

---

## 📄 License

MIT