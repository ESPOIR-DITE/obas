package academics

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/academics"
	"strconv"
	"testing"
)

func TestCreateSubject(t *testing.T) {
	obj := domain.Subject{"0000", "adp3", ""}
	result, err := CreateSubject(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestDeleteSubject(t *testing.T) {
	obj := domain.Subject{"0000", "adp3", ""}
	result, err := DeleteSubject(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetSubject(t *testing.T) {
	result, err := GetCourseSubject("0000", "")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetSubjects(t *testing.T) {
	result, err := GetSubjects()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestUpdateSubject(t *testing.T) {
	obj := domain.Subject{"0000", "adp3", ""}
	result, err := UpdateSubject(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}

/***want to read all the matric subject**/
func TestCreateSubject2(t *testing.T) {
	var subjects []domain.Subject
	courseSubject, err := GetAllCourseSubject()
	if err != nil {
		fmt.Println("do nothing")
	}
	for _, courseSubject := range courseSubject {
		if courseSubject.CourseId == "CCIT-EHMQR" {
			subject, err := GetSubject(courseSubject.SubjectId)
			if err != nil {
				fmt.Println("do nothing")
			}
			subjects = append(subjects, subject)
		}
	}
	writer(subjects)
	assert.NotNil(t, subjects)
}

//fmt.Println(" The Results", result)

func writer(subject []domain.Subject) {
	array_name := [3]string{"A", "B", "C"}

	f := excelize.NewFile()
	// Create a new sheet.
	//index := f.NewSheet("Sheet2")
	// Set value of a cell.
	for index, value1 := range subject {
		indexString := strconv.Itoa(index + 1)
		f.SetCellValue("Sheet1", array_name[0]+indexString, value1.Id)

		f.SetCellValue("Sheet1", array_name[1]+indexString, value1.Name)

		f.SetCellValue("Sheet1", array_name[2]+indexString, value1.SubjectDesc)

		fmt.Println(value1)

		//nuvo:=array_name[index]+indexString
		if err := f.SaveAs("C:/Users/Nicole Abrahams/go/src/obas/util/files/schoolSubject.xlsx"); err != nil {
			fmt.Println(err)
		}
	}
	//f.SetCellValue("Sheet1", "A1", "passionfruit")
	// Set active sheet of the workbook.
	//f.SetActiveSheet(index)
	// Save xlsx file by the given path.

}
