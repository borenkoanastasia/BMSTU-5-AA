package main

import ("fmt")

func main() {
	var str_start string = "abcsz"
	var max_len = 10
	var shifrXor string = generate_encryption_Xor_key(max_len)
	string_encryption_Cesar := get_string_encryption_Cesar(13)
	string_encryption_Xor := get_string_encryption_Xor(shifrXor)

	fmt.Println(str_start, "Цезарь", string_encryption_Cesar(str_start), "XOR", shifrXor, string_encryption_Xor(str_start), "Атбаш",
    string_encryption_Atbash(str_start))

    end := make(chan *queue_t, 120)
	conveer(end)
	fmt.Println("endd")
	var res *queue_t
	select{
		case res = <-end:
			fmt.Println("enddd")
			analysis(res);
	}
}
