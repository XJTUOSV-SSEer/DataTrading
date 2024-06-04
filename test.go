package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/clearmatics/bn256"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"log"
	"math/big"
	"os"
	"time"
)

// 指数运算测试
// 链上/链下哈希运算测试
//
//	func test() {
//		exp := big.NewInt(1)
//		rM := big.NewInt(232131)
//		rr := big.NewInt(456)
//		x1 := new(bn256.G1).ScalarBaseMult(exp)
//		x2 := new(bn256.G1).ScalarMult(x1, rM)
//		y1 := new(bn256.G2).ScalarBaseMult(rr)
//		// print(x.String())
//		// println(bn256.P)
//		// crypto.Keccak256()
//	}

func polynomial(a []*big.Int, secret *big.Int, i int, t int) (p_i *big.Int) {
	p_i = big.NewInt(secret.Int64())
	temA := big.NewInt(1)
	temB := big.NewInt(0)
	for j := 0; j < t; j = j + 1 {
		temB.Mul(a[j], temA)
		p_i.Add(p_i, temB)
		temA.Mul(temA, big.NewInt(int64(i)))
	}
	return p_i
}

func pvss_steup() (*big.Int, *bn256.G2) {
	psk, _ := rand.Int(rand.Reader, bn256.Order)
	ppk := new(bn256.G2).ScalarBaseMult(psk)
	return psk, ppk
}

func pvss_share(secret *big.Int, ppks []*bn256.G2, t int, n int) ([]*bn256.G2, []*bn256.G2, *bn256.G1, []*bn256.G1, []*bn256.GT, []*bn256.G2, []*big.Int, []*big.Int) {
	a := make([]*big.Int, t)
	E1 := make([]*bn256.G2, n)
	E2 := make([]*bn256.G2, n)
	A := make([]*bn256.G1, t)
	Hs := new(bn256.G1).ScalarBaseMult(secret)
	ck1s := make([]*big.Int, n)
	ck2s := make([]*big.Int, n)
	Y1 := make([]*bn256.GT, n)
	Y2 := make([]*bn256.G2, n)
	for i := 0; i < t; i = i + 1 {
		a[i], _ = rand.Int(rand.Reader, bn256.Order)
		if i >= 1 {
			A[i] = new(bn256.G1).ScalarBaseMult(a[i])
		}
	}
	for i := 0; i < n; i = i + 1 {
		r_i, _ := rand.Int(rand.Reader, bn256.Order)
		p_i := polynomial(a, secret, i+1, t)
		tem := new(bn256.G1).Set(Hs)
		exp := big.NewInt(int64(i + 1))
		r_p, _ := rand.Int(rand.Reader, bn256.Order)
		for j := 1; j < t; j = j + 1 {
			A_j := A[j].ScalarMult(A[j], exp)
			tem.Add(tem, A_j)
			exp.Mul(exp, big.NewInt(int64(i+1)))
		}
		tem_2 := bn256.Pair(tem, new(bn256.G2).ScalarBaseMult(big.NewInt(1)))
		tem.ScalarMult(tem, r_p)
		Y1[i] = bn256.Pair(tem, new(bn256.G2).ScalarBaseMult(big.NewInt(1)))
		Y2[i] = new(bn256.G2).ScalarBaseMult(r_p)
		E1[i] = new(bn256.G2).ScalarBaseMult(r_i)
		E2[i] = new(bn256.G2).ScalarMult(ppks[i], r_i)
		E2[i].Add(E2[i], new(bn256.G2).ScalarBaseMult(p_i))

		var buffer bytes.Buffer
		buffer.Write(E1[i].Marshal())
		buffer.Write(tem_2.Marshal())
		buffer.Write(Y1[i].Marshal())
		buffer.Write(Y2[i].Marshal())
		tem_b := crypto.Keccak256(buffer.Bytes())
		c := int64(binary.BigEndian.Uint64(tem_b[:]))
		ck1s[i] = big.NewInt(r_p.Int64() - c)
		ck2s[i] = big.NewInt(r_p.Int64() - c*r_i.Int64())
	}

	return E1, E2, Hs, A, Y1, Y2, ck1s, ck2s
}

func pvss_verify(E1 *bn256.G2, E2 *bn256.G2, Hs *bn256.G1, A []*bn256.G1, Y1 *bn256.GT, Y2 *bn256.G2, ck1s *big.Int, ck2s *big.Int, ppks *bn256.G2, i int) bool {
	flag := true
	exp := i
	tem := new(bn256.G1).Set(Hs)
	for j := 1; j < len(A); j = j + 1 {
		tt := new(bn256.G1).ScalarBaseMult(big.NewInt(int64(exp)))
		exp = exp * i
		tem.Add(tem, tt)
	}
	tem_GT := bn256.Pair(tem, new(bn256.G2).ScalarBaseMult(big.NewInt(1)))
	var buffer bytes.Buffer
	buffer.Write(E1.Marshal())
	buffer.Write(tem_GT.Marshal())
	buffer.Write(Y1.Marshal())
	buffer.Write(Y2.Marshal())
	tem_b := crypto.Keccak256(buffer.Bytes())
	c := int64(binary.BigEndian.Uint64(tem_b[:]))

	GT := new(bn256.GT).ScalarMult(tem_GT, big.NewInt(c))
	GT.Add(GT, bn256.Pair(new(bn256.G1).ScalarBaseMult(ck1s), E2))
	ppks_neg := new(bn256.G2).Set(ppks)
	ppks_neg.Neg(ppks_neg)
	GT.Add(GT, bn256.Pair(new(bn256.G1).ScalarBaseMult(ck2s), ppks_neg))
	if GT.String() != Y1.String() {
		flag = false
	}
	G2 := new(bn256.G2).ScalarMult(E1, big.NewInt(c))
	G2.Add(G2, new(bn256.G2).ScalarBaseMult(ck2s))
	if G2.String() != Y2.String() {
		flag = false
	}
	return flag
}

func pvss_recovery(x []int, ee []*bn256.G2) *bn256.G2 {
	secret := new(bn256.G2).ScalarBaseMult(big.NewInt(1))
	for i := 0; i < len(x); i = i + 1 {
		lambda_i := big.NewInt(1)
		for j := 0; j < len(x); j = j + 1 {
			if i == j {
				continue
			}
			tem := big.NewInt(int64(j))
			tem.Div(tem, big.NewInt(int64(j-i)))
			lambda_i.Mul(lambda_i, tem)
		}
		secret.Add(secret, new(bn256.G2).ScalarMult(ee[i], lambda_i))
	}
	return secret
}

func mps_keygen(n int) ([][]*big.Int, [][]*bn256.G2, []*bn256.G2) {
	sks := make([][]*big.Int, n)
	pks := make([][]*bn256.G2, n)
	X_A := new(bn256.G2).ScalarBaseMult(big.NewInt(1))
	Y_A_1 := new(bn256.G2).ScalarBaseMult(big.NewInt(1))
	Y_A_2 := new(bn256.G2).ScalarBaseMult(big.NewInt(1))
	for i := 0; i < n; i = i + 1 {
		sks[i] = make([]*big.Int, 3)
		pks[i] = make([]*bn256.G2, 3)
		sks[i][0], _ = rand.Int(rand.Reader, bn256.Order)
		sks[i][1], _ = rand.Int(rand.Reader, bn256.Order)
		sks[i][2], _ = rand.Int(rand.Reader, bn256.Order)
		pks[i][0] = new(bn256.G2).ScalarBaseMult(sks[i][0])
		pks[i][1] = new(bn256.G2).ScalarBaseMult(sks[i][1])
		pks[i][2] = new(bn256.G2).ScalarBaseMult(sks[i][2])
		tem_hash := new(bn256.G2).Add(pks[i][0], pks[i][1])
		tem_hash.Add(tem_hash, pks[i][2])
		tem_byte := md5.Sum(tem_hash.Marshal())
		r_i := big.NewInt(int64(binary.BigEndian.Uint64(tem_byte[:])))
		X_A.Add(X_A, new(bn256.G2).ScalarMult(pks[i][0], r_i))
		Y_A_1.Add(Y_A_1, new(bn256.G2).ScalarMult(pks[i][1], r_i))
		Y_A_2.Add(Y_A_2, new(bn256.G2).ScalarMult(pks[i][2], r_i))
	}
	vk := make([]*bn256.G2, 3)
	vk[0] = X_A
	vk[1] = Y_A_1
	vk[2] = Y_A_2
	return sks, pks, vk
}

func mps_sign(m *big.Int, sk []*big.Int) (*big.Int, *bn256.G1, *bn256.G1) {
	mm := md5.Sum(m.Bytes())
	mm_r := big.NewInt(int64(binary.BigEndian.Uint64(mm[:])))
	h := new(bn256.G1).ScalarBaseMult(mm_r)
	tem := big.NewInt(sk[0].Int64() + m.Int64()*sk[1].Int64() + mm_r.Int64()*sk[2].Int64())
	pi_i_2 := new(bn256.G1).ScalarMult(h, tem)
	return mm_r, h, pi_i_2
}

func mps_aggregate(pks [][]*bn256.G2, pi_i_2_s []*bn256.G1) *bn256.G1 {
	pi_A_2 := new(bn256.G1).ScalarBaseMult(big.NewInt(1))
	for i := 0; i < len(pi_i_2_s); i = i + 1 {
		tem_hash := new(bn256.G2).Add(pks[i][0], pks[i][1])
		tem_hash.Add(tem_hash, pks[i][2])
		tem_byte := md5.Sum(tem_hash.Marshal())
		r_i := big.NewInt(int64(binary.BigEndian.Uint64(tem_byte[:])))
		pi_A_2.Add(pi_A_2, new(bn256.G1).ScalarMult(pi_i_2_s[i], r_i))
	}
	return pi_A_2
}
func mps_verify(h *bn256.G1, pi_A_2 *bn256.G1, mm_r *big.Int, m *big.Int, vk []*bn256.G2) bool {
	temG2 := new(bn256.G2).Set(vk[0])
	temG2.Add(temG2, new(bn256.G2).ScalarMult(vk[1], m))
	temG2.Add(temG2, new(bn256.G2).ScalarMult(vk[2], mm_r))
	t1 := bn256.Pair(h, temG2)
	t2 := bn256.Pair(pi_A_2, new(bn256.G2).ScalarBaseMult(big.NewInt(1)))
	return t1.String() == t2.String()
}

func Test_AES(size int) {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	if size > len(file) {
		println("文件长度不足")
		return
	}
	file = file[:size]
	origData := file                  // 待加密的数据
	key := []byte("9876787656785679") // 加密的密钥
	// log.Println("原文：", string(origData))

	log.Println("------------------ CFB模式 --------------------")
	start := time.Now()
	encrypted := AesEncryptCFB(origData, key)
	end := time.Now()
	println("加密尺寸为", size, "大小的文件耗时", end.Sub(start).Microseconds())
	// log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	// log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	start = time.Now()
	AesDecryptCFB(encrypted, key)
	end = time.Now()
	println("解密尺寸为", size, "大小的文件耗时", end.Sub(start).Microseconds())
	// log.Println("解密结果：", string(decrypted))
}
func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}
func AesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}

func construct_pi_fs(h *bn256.G1, pi_A_2 *bn256.G1, m_f []byte, sks int64, mm int64) (*bn256.G1, *bn256.G1, *bn256.G1, int64, int64, int64) {
	r, _ := rand.Int(rand.Reader, bn256.Order)
	pi_1 := new(bn256.G1).ScalarMult(h, r)
	pi_2 := new(bn256.G1).ScalarMult(pi_A_2, r)
	r_m, _ := rand.Int(rand.Reader, bn256.Order)
	pi_2_rm_neg := new(bn256.G1).ScalarMult(pi_2, r_m)
	pi_2_rm_neg.Neg(pi_2_rm_neg)
	// g2_rm := new(bn256.G2).ScalarBaseMult(r_m)

	var buffer bytes.Buffer
	buffer.Write(m_f)
	buffer.Write(pi_2_rm_neg.Marshal())
	// buffer.Write(g2_rm.Marshal())
	buffer.Write(pi_2.Marshal())
	buffer.Write(new(bn256.G2).ScalarBaseMult(big.NewInt(1)).Marshal())
	tem_b := crypto.Keccak256(buffer.Bytes())
	c := int64(binary.BigEndian.Uint64(tem_b[:]))
	ck1 := r_m.Int64() - c

	ck2 := sks * (r_m.Int64() - c)
	ck3 := mm * (r_m.Int64() - c)
	return pi_1, pi_2, pi_2_rm_neg, ck1, ck2, ck3
}

// pks,pkt,gc,h3_idf,e1,ee1,e2,ee2,tl,ttl
// m_ok
// r_f,rr_f,Ks,r_l
// Y_1, Y_2, Y_3, Y_4, Y_5, ck1, ck2, ck3, ck4
func construct_pi_ok(pks *bn256.G1, pkt *bn256.G1, gc *bn256.G1, h3_idf *bn256.G1, e1 *bn256.G1, ee1 *bn256.G1, e2 *bn256.G1, ee2 *bn256.G1, tl *bn256.G1, ttl *bn256.G1, m_ok []byte, r_f int64, r_ff int64, Ks int64, r_l int64) (*bn256.G1, *bn256.G1, *bn256.G1, *bn256.G1, *bn256.G1, int64, int64, int64, int64) {
	r_ok, _ := rand.Int(rand.Reader, bn256.Order)
	Y_1 := new(bn256.G1).ScalarBaseMult(r_ok)
	Y_2 := new(bn256.G1).ScalarMult(pks, r_ok)
	Y_2.Add(Y_2, Y_1)
	Y_3 := new(bn256.G1).ScalarMult(pkt, r_ok)
	Y_3.Add(Y_3, Y_1)
	Y_4 := new(bn256.G1).ScalarMult(gc, r_ok)
	Y_5 := new(bn256.G1).ScalarMult(h3_idf, r_ok)
	var buffer bytes.Buffer
	buffer.Write(Y_1.Marshal())
	buffer.Write(Y_2.Marshal())
	buffer.Write(Y_3.Marshal())
	buffer.Write(Y_4.Marshal())
	buffer.Write(Y_5.Marshal())
	buffer.Write(e1.Marshal())
	buffer.Write(ee1.Marshal())
	buffer.Write(e2.Marshal())
	buffer.Write(ee2.Marshal())
	buffer.Write(tl.Marshal())
	buffer.Write(ttl.Marshal())
	buffer.Write(m_ok)
	tem_b := crypto.Keccak256(buffer.Bytes())
	c := int64(binary.BigEndian.Uint64(tem_b[:]))
	ck1 := r_ok.Int64() - r_f*c
	ck2 := r_ok.Int64() - r_ff*c
	ck3 := r_ok.Int64() - Ks*c
	ck4 := r_ok.Int64() - c*r_l
	return Y_1, Y_2, Y_3, Y_4, Y_5, ck1, ck2, ck3, ck4
}

func constrct_pi_fj(tid *big.Int, ek *big.Int, m *big.Int) (*big.Int, *big.Int, *big.Int) {
	r_p1, _ := rand.Int(rand.Reader, bn256.Order)
	r_p2, _ := rand.Int(rand.Reader, bn256.Order)
	r_p3, _ := rand.Int(rand.Reader, bn256.Order)
	p1 := big.NewInt(0)
	p2 := big.NewInt(0)
	p3 := big.NewInt(0)
	p1.Mul(r_p1, tid)
	p2.Mul(r_p2, ek)
	p3.Mul(r_p3, m)
	return p1, p2, p3
}

func main() {
	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Panic("failed to Dail", err)
	}
	defer conn.Close()

	demo, err := NewTest(common.HexToAddress("0xf595cB47B06601b69E413a84dA70c25e41c6C099"), conn)

	if err != nil {
		log.Panic("failed to New", err)
	}
	privateKeyStr := "9e72e5257645bebc6e3423696be498c6973cc23cee4aaad507d04331d51fcef6"
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = nil
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = big.NewInt(0)
	//var start,end

	// PVSS.Steup耗时测试
	n := 10 // PVSS阶段的参与者
	start := time.Now()
	psks := make([]*big.Int, n)
	ppks := make([]*bn256.G2, n)
	for i := 0; i < n; i = i + 1 {
		psks[i], ppks[i] = pvss_steup()
	}
	end := time.Now()
	println("初始化", n, "个PVSS.STEUP过程耗时", end.Sub(start).Microseconds())

	// pi_fj构建过程耗时
	start = time.Now()
	tid, _ := rand.Int(rand.Reader, bn256.Order)
	ek, _ := rand.Int(rand.Reader, bn256.Order)
	mmm, _ := rand.Int(rand.Reader, bn256.Order)
	p1, p2, p3 := constrct_pi_fj(tid, ek, mmm)
	c1 := big.NewInt(0)
	c2 := big.NewInt(0)
	c1.Mul(tid, ek)
	c2.Mul(mmm, ek)
	end = time.Now()
	println("pi_fj构建过程耗时", end.Sub(start).Microseconds())

	// pi_fj验证过程耗时
	start = time.Now()
	_, err = demo.VerifyPiFj(auth, p1, p2, p3, c1, c2)
	end = time.Now()
	println("pi_fj验证过程耗时", end.Sub(start).Milliseconds())

	return

	// PVSS.share耗时测试
	start = time.Now()
	t := n // 秘密恢复数量
	//E1 := make([]*bn256.G2, n)
	//E2 := make([]*bn256.G2, n)
	E1, E2, Hs, A, Y1, Y2, ck1s, ck2s := pvss_share(big.NewInt(12345), ppks, t, n)
	end = time.Now()
	println("分享", t, "个PVSS.SHARE过程耗时", end.Sub(start).Microseconds())

	// PVSS.verify耗时测试
	for i := 0; i < n; i = i + 1 {
		start = time.Now()
		pvss_verify(E1[i], E2[i], Hs, A, Y1[i], Y2[i], ck1s[i], ck2s[i], ppks[i], i+1)
		end = time.Now()
		println("PVSS.Verify过程耗时", end.Sub(start).Microseconds())
	}

	// PVSS.recovery耗时测试
	start = time.Now()
	var x []int
	for i := 0; i < t; i = i + 1 {
		x = append(x, i+1)
	}
	ee := make([]*bn256.G2, t)
	for i := 0; i < t; i = i + 1 {
		E1[i].ScalarMult(E1[i], psks[i])
		E1[i].Neg(E1[i])
		ee[i] = new(bn256.G2)
		ee[i].Add(E1[i], E2[i])
	}
	pvss_recovery(x, ee)
	end = time.Now()
	println(t, "个分片执行PVSS.Recover过程耗时", end.Sub(start).Microseconds())

	// MPS.KeyGen耗时测试
	start = time.Now()
	nn := 10
	sks, pks, vk := mps_keygen(nn)
	end = time.Now()
	println("密钥生成", nn, "个MPS.KeyGen过程耗时", end.Sub(start).Microseconds())

	// MPS.Sign耗时测试
	start = time.Now()
	pi_i_2_s := make([]*bn256.G1, nn)
	mm_r := new(big.Int)
	h := new(bn256.G1)
	m := big.NewInt(12345)
	for i := 0; i < nn; i = i + 1 {
		mm_r, h, pi_i_2_s[i] = mps_sign(m, sks[i])
	}
	end = time.Now()
	println("签名", nn, "个MPS.Sign过程耗时", end.Sub(start).Microseconds())

	// MPS.Aggergate耗时测试
	start = time.Now()
	pi_A_2 := mps_aggregate(pks, pi_i_2_s)
	end = time.Now()
	println("聚合", nn, "个MPS签名过程耗时", end.Sub(start).Microseconds())

	// MPS.Verify耗时
	start = time.Now()
	mps_verify(h, pi_A_2, mm_r, m, vk)
	end = time.Now()
	println("验证MPS签名过程耗时", end.Sub(start).Microseconds())

	// pi_fs构建耗时
	start = time.Now()
	pi_1, pi_2, pi_2_rm_neg, ck1, ck2, ck3 := construct_pi_fs(h, pi_A_2, big.NewInt(12345).Bytes(), 12345, mm_r.Int64())

	end = time.Now()
	println("pi_fs构建过程耗时", end.Sub(start).Microseconds())

	//exp := big.NewInt(1)
	//rM := big.NewInt(3)
	//rr := big.NewInt(1)
	//x1 := new(bn256.G1).ScalarBaseMult(exp)
	//x2 := new(bn256.G1).ScalarMult(x1, rM)
	//x2.Neg(x2)
	//x3 := new(bn256.G1).ScalarBaseMult(rM)
	//flag := new(int)
	//*flag = 0
	//for i := 0; i < 64; i = i + 1 {
	//	if x2.Marshal()[i] != x3.Marshal()[i] {
	//		*flag = 1
	//	}
	//}
	//println(*flag)
	// y1 := new(bn256.G2).ScalarBaseMult(rr)
	// y1.Neg(y1)

	// pi_fs验证耗时
	start = time.Now()
	g1 := WritepiFsG1{Pi1: pi_1.Marshal(), Pi2: pi_2.Marshal(), Pi2RmNeg: pi_2_rm_neg.Marshal()}
	ck := WritepiFsCk{Ck1: big.NewInt(ck1), Ck2: big.NewInt(ck2), Ck3: big.NewInt(ck3)}
	g2 := WritepiFsG2{XA: vk[0].Marshal(), YA1: vk[1].Marshal(), YA2: vk[2].Marshal(), G2: new(bn256.G2).ScalarBaseMult(big.NewInt(1)).Marshal()}
	tx, err := demo.VerifyPiFs(auth, g1, g2, ck, big.NewInt(12345).Bytes())
	// println(g1, ck, g2)
	end = time.Now()
	if err != nil {
		log.Panic("failed to set ", err)
	}
	fmt.Println("pi_fs gas_cost:", tx.Cost())
	println("pi_fs验证过程耗时", end.Sub(start).Milliseconds())

	// pi_ok构建过程耗时
	start = time.Now()
	_, pks_t, _ := bn256.RandomG1(rand.Reader)
	_, pkt, _ := bn256.RandomG1(rand.Reader)
	_, gc, _ := bn256.RandomG1(rand.Reader)
	_, h3_idf, _ := bn256.RandomG1(rand.Reader)
	_, e1, _ := bn256.RandomG1(rand.Reader)
	_, ee1, _ := bn256.RandomG1(rand.Reader)
	_, e2, _ := bn256.RandomG1(rand.Reader)
	_, ee2, _ := bn256.RandomG1(rand.Reader)
	_, tl, _ := bn256.RandomG1(rand.Reader)
	_, ttl, _ := bn256.RandomG1(rand.Reader)
	m_ok := big.NewInt(12345).Bytes()
	r_f, _ := rand.Int(rand.Reader, bn256.Order)
	rr_f, _ := rand.Int(rand.Reader, bn256.Order)
	Ks, _ := rand.Int(rand.Reader, bn256.Order)
	r_l, _ := rand.Int(rand.Reader, bn256.Order)
	Y_1, Y_2, Y_3, Y_4, Y_5, ck11, ck22, ck33, ck44 := construct_pi_ok(pks_t, pkt, gc, h3_idf, e1, ee1, e2, ee2, tl, ttl, m_ok, r_f.Int64(), rr_f.Int64(), Ks.Int64(), r_l.Int64())
	end = time.Now()
	println("pi_ok构建过程耗时", end.Sub(start).Microseconds())

	// pi_ok验证过程耗时
	start = time.Now()
	Y := WriteY{Y1: Y_1.Marshal(), Y2: Y_2.Marshal(), Y3: Y_3.Marshal(), Y4: Y_4.Marshal(), Y5: Y_5.Marshal()}
	ckk := WritepiOkCk{Ck1: big.NewInt(ck11), Ck2: big.NewInt(ck22), Ck3: big.NewInt(ck33), Ck4: big.NewInt(ck44)}
	tt := WritepiOkTt{E1: e1.Marshal(), Ee1: ee1.Marshal(), E2: e2.Marshal(), Ee2: ee2.Marshal(), Tl: tl.Marshal(), Ttl: ttl.Marshal()}
	ttx, err := demo.VerifyPiOk(auth, Y, tt, m_ok, ckk, new(bn256.G1).ScalarBaseMult(big.NewInt(1)).Marshal(), pks_t.Marshal(), pkt.Marshal(), gc.Marshal(), h3_idf.Marshal())
	// println(Y, ckk, tt)
	end = time.Now()
	if err != nil {
		log.Panic("failed to set ", err)
	}
	fmt.Println("pi_ok cost:", ttx.GasFeeCap())
	println("pi_ok验证过程耗时", end.Sub(start).Milliseconds())

	// abigen --abi test.abi --pkg main --type test --out testabi.go
	//println(y1.Marshal())
	//tx, err := demo.Test(auth, x1.Marshal(), x2.Marshal(), rM, y1.Marshal())
	//if err != nil {
	//	log.Panic("failed to set ", err)
	//}
	//fmt.Println("setmsg tx:", tx.Hash())
	//res, _ := demo.GetCode(nil)
	//println(res.String())
	// Test_AES(5120000)

}
