package main

import (
	"encoding/json"
	"fmt"
)

var strJson = `{
    "m3ua" : {
         "opc" : 2057,
         "dpc" : 4114,
         "rc" : 1,
         "ni" : 0,
         "na" : 1
    },
 
    "sccp" : {
        "cldpn": {
                "gti" : 18,
                "spc" : 0,
                "ssn" : 6,
                "translationType" : 10,
                "numberingPlan" : 1,
                "encodingScheme" : 1,
                "natureOfAddressIndicator" : 4,
                "digits" : "38063210599"
        },
        "clgpn" : {
                "gti" : 18,
                "spc" : 0,
                "ssn" : 7,
                "translationType" : 10,
                "numberingPlan" : 1,
                "encodingScheme" : 2,
                "natureOfAddressIndicator": 4,
                "digits" : "14485000002"
        }
    },
 
    "tcap" : {
        "invokeid" : 1,
        "localValue" : 60,
        "destRef" : "96202019900500511",
        "origRef" : "91947105410301"
    },
    "la": "10.10.220.84:2905",
    "ra": "10.10.220.105:2905",
    "flow": [["send", "tc_begin", "unstructured_notify", "", "0"], ["recieve", "tc_continue", "return_error", "", "0"], ["send", "tc_end", "0", "", "1"]],
    "update_data": {
        "tc_end" : {
            "sccp" : {
                "cldpn": {
                        "numberingPlan" : 4
                },
                "clgpn" : {
                        "natureOfAddressIndicator": 4
                }
            }
        },
        "tc_continue" : {
            "sccp" : {
                "cldpn": {
                        "numberingPlan" : 4
                },
                "clgpn" : {
                        "natureOfAddressIndicator": 4
                }
            }
        }
    }
}
 `

func main() {
	fmt.Println("Unknow format json")

	data := make(map[string]interface{})

	err := json.Unmarshal([]byte(strJson), &data)
	if err != nil {
		fmt.Println("Received error while Unmarshling", err)
	} else {
		fmt.Println(data)
	}

	for key, val := range data {
		switch val.(type) {
		case int:
			fmt.Println(key, "int")
		case float32, float64:
			fmt.Println(key, "float")

		case string:
			fmt.Println(key, "string")

		case []string:
			fmt.Println(key, "[]string")

		case [][]string:
			fmt.Println(key, "[][5]string")

		case map[string]interface{}:
			fmt.Println(key, "json object")
		default:
			fmt.Println(key, "unknow")
		}
	}
}
