package sutfinallab

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestNotblankValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("name not blank", func(t *testing.T) {
		e := Employee{
			Name:       "",
			Email:      "mamoss@gmail.com",
			EmployeeID: "M88888888",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("success all"))
	})
}
