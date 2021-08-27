package main

import "testing"

func TestGenerateKey(t *testing.T) {

	testCases := []struct {
		shift   int
		inputs  []rune
		outputs []rune
	}{
		{
			shift:   3,
			inputs:  []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
			outputs: []rune("DEFGHIJKLMNOPQRSTUVWXYZABC"),
		},
		{
			shift:   10,
			inputs:  []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
			outputs: []rune("KLMNOPQRSTUVWXYZABCDEFGHIJ"),
		},
		{
			shift:   25,
			inputs:  []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
			outputs: []rune("ZABCDEFGHIJKLMNOPQRSTUVWXY"),
		},
	}

	for _, tC := range testCases {

		if len(tC.inputs) != len(tC.outputs) {

			t.Fatalf(
				"Error in test case: input length must match output length. shift: %d. inputs: %s(%d). outputs: %s(%d).",
				tC.shift,
				string(tC.inputs),
				len(tC.inputs),
				string(tC.outputs),
				len(tC.outputs),
			)

		}

		key := generateKey(tC.shift)

		for i := 0; i < len(tC.inputs); i += 1 {

			if key[tC.inputs[i]] != tC.outputs[i] {

				t.Errorf(
					"Key mapping is incorrect. Shift: %d. Expected: %c => %c. Got: %c => %c",
					tC.shift,
					tC.inputs[i],
					tC.outputs[i],
					tC.inputs[i],
					key[tC.inputs[i]],
				)

			}

		}

	}

}

func TestEncrypt(t *testing.T) {

	testCases := []struct {
		msg       string
		encrypted string
	}{
		{
			msg:       "",
			encrypted: "",
		},
		{
			msg:       "Hello, World!",
			encrypted: "KHOORZRUOG",
		},
		{
			msg:       "H e l l oWoRlD",
			encrypted: "KHOORZRUOG",
		},
		{
			msg:       "??\"H\"ello?World!'\"",
			encrypted: "KHOORZRUOG",
		},
		{
			msg:       "\033[31;1;4mHello World\033[0m",
			encrypted: "KHOORZRUOG",
		},
	}

	for _, tC := range testCases {

		key := generateKey(3)
		encrypted := encrypt(key, tC.msg)

		if encrypted != tC.encrypted {
			t.Errorf("Encrypted value is incorrect. msg: %s. Expected: %s. Got: %s.", tC.msg, tC.encrypted, encrypted)
		}

	}

}
