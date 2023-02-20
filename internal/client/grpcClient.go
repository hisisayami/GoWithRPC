package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/go-inventory-grpc/internal/endpoint"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("didn't connect to port 9000: %v", err)
	}

	defer conn.Close()

	c := endpoint.NewInventoryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.Register(ctx, &endpoint.Message{})
	if err != nil {
		log.Fatalf("didn't connect to port 9000: %v", err)
	}
	log.Println("response :", r)

	//CreateStaff
	log.Printf("------------------CreateStaff RPC--------------------, \n")

	createStaff, err := c.CreateStaff(context.Background(), &endpoint.CreateStaffRequest{
		Name:  "cherry",
		Email: "cherry27@gmail.com",
	})
	if err != nil {
		log.Fatalf("failed to create staff: %v", err)
	}

	log.Println(createStaff.Name, createStaff.Email, createStaff.Id)

	//GetStaffById
	log.Printf("--------------------GetStaffById RPC----------------, \n")

	getStaffById, err := c.GetStaffById(context.Background(), &endpoint.GetStaffByIdRequest{
		Id: 8,
	})
	if err != nil {
		log.Fatalf("failed to get staff by is: %v", err)
	}
	log.Println(getStaffById.Name, getStaffById.Email, getStaffById.Id)

	//DeleteStaffById
	// log.Printf("-------------------DeleteStaffById RPC-------------, \n")

	// _, err = c.DeleteStaffById(ctx, &endpoint.DeleteStaffByIdRequest{
	// 	Id: 10,
	// })
	// if err != nil {
	// 	log.Fatalf("failed to delete staff by id: %v", err)
	// }

	//UpdateStaffById
	log.Printf("-------------------UpdateStaffById RPC-------------, \n")

	updatedStaff, err := c.UpdateStaffById(ctx, &endpoint.UpdateStaffByIdRequest{
		Id:    4,
		Name:  "pijan",
		Email: "pijan@gamaa.com",
	})
	if err != nil {
		log.Fatalf("failed to update staff by id: %v", err)
	}

	fmt.Println("Staff updated!!", updatedStaff)

	getAllStaff, err := c.GetAllStaff(ctx, &endpoint.GetAllStaffRequest{})
	if err != nil {
		log.Fatalf("failed to get list of staffs: %v", err)
	}

	fmt.Println("Staff List!!", getAllStaff)

	createCategory, err := c.CreateCategory(context.Background(), &endpoint.CreateCategoryRequest{
		CategoryName:        "snacks",
		CategoryDescription: "samosa",
	})
	if err != nil {
		log.Fatalf("failed to create category: %v", err)
	}
	fmt.Println("Category created!!", createCategory)

	createProduct, err := c.CreateProduct(context.Background(), &endpoint.CreateProductRequest{
		ProductName:        "chowmine",
		ProductDescription: "vegetables chowmine",
		ProductQuantity:    9,
		UnitPrice:          100,
	})
	if err != nil {
		log.Fatalf("failed to create product: %v", err)
	}
	fmt.Println("Product created!!", createProduct)

}

// func main() {
// 	type address struct {
// 		name    string
// 		address string
// 	}

// 	staff1 := address{
// 		name:    "pijan",
// 		address: "zoo",
// 	}

// 	staff2 := address{
// 		name:    "khijan",
// 		address: "jungle",
// 	}

// 	staff3 := address{
// 		name:    "Sijan",
// 		address: "forest",
// 	}

// 	type staffDetails struct {
// 		name    string
// 		phone   int
// 		address string
// 	}

// 	staff4 := staffDetails{
// 		name:  "pijan",
// 		phone: 0,
// 	}

// 	staff5 := staffDetails{
// 		name:  "khijan",
// 		phone: 1,
// 	}

// 	staff6 := staffDetails{
// 		name:  "Sijan",
// 		phone: 2,
// 	}

// 	staffAddress := []address{staff1, staff2, staff3}
// 	staffDetail := []staffDetails{staff4, staff5, staff6}

// 	//strings := []string{"hello", "world", "ppp", "ccc", "hyahyahya"}

// 	for i, s := range staffAddress {
// 		for _, k := range staffDetail {
// 			// add := staffDetails{
// 			// 	name:    k.name,
// 			// 	phone:   k.phone,
// 			// 	address: s.address,
// 			//}

// 			//fmt.Println("j--k--", j, k)
// 			//staffDetail = append(staffDetail, add)
// 			//	fmt.Println("---add----", add)
// 			if s.name == k.name {
// 				k.address = s.address
// 			}
// 			fmt.Println("k-------", k)

// 		}

// 		fmt.Println(i, s)

// 	}

// 	//fmt.Println(staffAddress)
// 	fmt.Println("=====Staff Details=====", staffDetail)

// }
