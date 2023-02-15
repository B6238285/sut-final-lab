package sutfinallab

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid:"required~success all"`
	Email      string
	EmployeeID string `valid: "matches(([JML]//d{8}))$~not match employee id"`
}

func TestEmployeeValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("success all", func(t *testing.T) {
		e := Employee{
			Name:       "moss",
			Email:      "mamoss@gmail.com",
			EmployeeID: "M88888888",
		}

		ok, err := govalidator.ValidateStruct(e)
		g.Expect(ok).To(gomega.BeTrue())
		g.Expect(err).To(gomega.BeNil())
	})
}
