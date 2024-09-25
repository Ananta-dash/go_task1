package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "time"
    "os"
)

type Resume struct {
    Name           string   `json:"name"`
    Email          string   `json:"email"`
    Phone          string   `json:"phone"`
    Experience     []string `json:"experience"`
    Skills         []string `json:"skills"`
    Education      string   `json:"education"`
    Certifications []string `json:"certifications"`
}

func randomName() string {
    firstNames := []string{"John", "Jane", "Alex", "Emily", "Chris", "Katie"}
    lastNames := []string{"Doe", "Smith", "Johnson", "Brown", "Davis", "Martinez"}
    rand.Seed(time.Now().UnixNano())
    firstName := firstNames[rand.Intn(len(firstNames))]
    lastName := lastNames[rand.Intn(len(lastNames))]
    return firstName + " " + lastName
}

func randomEmail(name string) string {
    domains := []string{"example.com", "email.com", "test.com"}
    rand.Seed(time.Now().UnixNano())
    return name + "@" + domains[rand.Intn(len(domains))]
}

func randomPhone() string {
    rand.Seed(time.Now().UnixNano())
    return fmt.Sprintf("(%03d) %03d-%04d", rand.Intn(1000), rand.Intn(1000), rand.Intn(10000))
}

func randomExperience() []string {
    experiences := [][]string{
        {"Software Engineer at ABC Corp", "Frontend Developer at XYZ Inc."},
        {"Backend Developer at 123 Tech", "Systems Analyst at DEF Ltd."},
    }
    rand.Seed(time.Now().UnixNano())
    return experiences[rand.Intn(len(experiences))]
}

func randomSkills() []string {
    skills := []string{"Go", "Python", "AWS", "Docker", "Kubernetes", "ReactJS", "PostgreSQL"}
    rand.Seed(time.Now().UnixNano())
    numSkills := rand.Intn(4) + 3 // Generate between 3-6 skills
    selectedSkills := make([]string, numSkills)
    for i := 0; i < numSkills; i++ {
        selectedSkills[i] = skills[rand.Intn(len(skills))]
    }
    return selectedSkills
}

func generateResume() Resume {
    name := randomName()
    email := randomEmail(name)
    phone := randomPhone()
    experience := randomExperience()
    skills := randomSkills()
    education := "BS in CS"
    certifications := []string{"AWS Certified Solutions Architect", "Certified Kubernetes Administrator"}

    return Resume{
        Name:           name,
        Email:          email,
        Phone:          phone,
        Experience:     experience,
        Skills:         skills,
        Education:      education,
        Certifications: certifications,
    }
}

func main() {
    resume := generateResume()
    resumeJSON, _ := json.MarshalIndent(resume, "", "  ")
    err := os.WriteFile("resume.json", resumeJSON, 0644)
    if err != nil {
        fmt.Println("Task Unsuccessful - writing to file:", err)
    } else {
        fmt.Println("Task Successful - Resume saved to resume.json")
    }
}
