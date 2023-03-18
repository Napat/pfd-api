# PFD: Pie Fire Dire  

กำหนด beef list api datasource อยู่ที่  
-> [https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text](https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text)

ซึ่งจะเป็นรายชื่อของเนื้อหลายชนิดปะปนกันอยู่ ตัวอย่างเช่น

```Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.```

ทุกคำเป็นชื่อชนิดเนื้อทั้งหมด ซึ่งจะมีตัวอักษร `,`, `.` รวมถึงช่องว่างเช่น space ปะปนอยู่  

## Application API specification

- Beef API endpoint: `/beef/summary`

--- sample response --

```json
{
    "beef": {
        "t-bone": 4,
        "fatback": 1,
        "pastrami": 1,
        "pork": 1,
        "meatloaf": 1,
        "jowl": 1,
        "enim": 1,
        "bresaola": 1
    }
}
```

## How to demonstrate

``` sh
go test -race ./...
go run cmd/main.go
curl http://localhost:8000/beef/summary
```
