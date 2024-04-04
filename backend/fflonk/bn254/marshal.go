package fflonk

import (
	curve "github.com/consensys/gnark-crypto/ecc/bn254"

	"io"

	"github.com/consensys/gnark-crypto/ecc/bn254/kzg"
)

// WriteRawTo writes binary encoding of Proof to w without point compression
func (proof *Proof) WriteRawTo(w io.Writer) (int64, error) {
	return proof.writeTo(w, curve.RawEncoding())
}

// WriteTo writes binary encoding of Proof to w with point compression
func (proof *Proof) WriteTo(w io.Writer) (int64, error) {
	return proof.writeTo(w)
}

func (proof *Proof) writeTo(w io.Writer, options ...func(*curve.Encoder)) (int64, error) {
	enc := curve.NewEncoder(w, options...)

	toEncode := []interface{}{
		&proof.LROEntangled,
		&proof.Z,
		&proof.ZEntangled,
		&proof.HEntangled,
		proof.BsbComEntangled,
		&proof.BatchOpeningProof.SOpeningProof.W,
		&proof.BatchOpeningProof.SOpeningProof.WPrime,
		proof.BatchOpeningProof.SOpeningProof.ClaimedValues,
		proof.BatchOpeningProof.ClaimedValues,
	}

	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			return enc.BytesWritten(), err
		}
	}

	return enc.BytesWritten(), nil
}

// ReadFrom reads binary representation of Proof from r
func (proof *Proof) ReadFrom(r io.Reader) (int64, error) {
	dec := curve.NewDecoder(r)
	toDecode := []interface{}{
		&proof.LROEntangled,
		&proof.ZEntangled,
		&proof.ZEntangled,
		&proof.HEntangled,
		&proof.BsbComEntangled,
		&proof.BatchOpeningProof.SOpeningProof.W,
		&proof.BatchOpeningProof.SOpeningProof.WPrime,
		&proof.BatchOpeningProof.SOpeningProof.ClaimedValues,
		&proof.BatchOpeningProof.ClaimedValues,
	}

	for _, v := range toDecode {
		if err := dec.Decode(v); err != nil {
			return dec.BytesRead(), err
		}
	}

	if proof.BsbComEntangled == nil {
		proof.BsbComEntangled = []kzg.Digest{}
	}

	return dec.BytesRead(), nil
}

// WriteTo writes binary encoding of ProvingKey to w
func (pk *ProvingKey) WriteTo(w io.Writer) (n int64, err error) {
	return pk.writeTo(w, true)
}

// WriteRawTo writes binary encoding of ProvingKey to w without point compression
func (pk *ProvingKey) WriteRawTo(w io.Writer) (n int64, err error) {
	return pk.writeTo(w, false)
}

func (pk *ProvingKey) writeTo(w io.Writer, withCompression bool) (n int64, err error) {
	// encode the verifying key
	if withCompression {
		n, err = pk.Vk.WriteTo(w)
	} else {
		n, err = pk.Vk.WriteRawTo(w)
	}
	if err != nil {
		return
	}

	var n2 int64
	// KZG key
	if withCompression {
		n2, err = pk.Kzg.WriteTo(w)
	} else {
		n2, err = pk.Kzg.WriteRawTo(w)
	}
	if err != nil {
		return
	}
	n += n2

	return n, nil
}

// ReadFrom reads from binary representation in r into ProvingKey
func (pk *ProvingKey) ReadFrom(r io.Reader) (int64, error) {
	return pk.readFrom(r, true)
}

// UnsafeReadFrom reads from binary representation in r into ProvingKey without subgroup checks
func (pk *ProvingKey) UnsafeReadFrom(r io.Reader) (int64, error) {
	return pk.readFrom(r, false)
}

func (pk *ProvingKey) readFrom(r io.Reader, withSubgroupChecks bool) (int64, error) {
	pk.Vk = &VerifyingKey{}
	n, err := pk.Vk.ReadFrom(r)
	if err != nil {
		return n, err
	}

	var n2 int64
	if withSubgroupChecks {
		n2, err = pk.Kzg.ReadFrom(r)
	} else {
		n2, err = pk.Kzg.UnsafeReadFrom(r)
	}
	n += n2
	if err != nil {
		return n, err
	}
	return n, err
}

// WriteTo writes binary encoding of VerifyingKey to w
func (vk *VerifyingKey) WriteTo(w io.Writer) (n int64, err error) {
	return vk.writeTo(w)
}

// WriteRawTo writes binary encoding of VerifyingKey to w without point compression
func (vk *VerifyingKey) WriteRawTo(w io.Writer) (int64, error) {
	return vk.writeTo(w, curve.RawEncoding())
}

func (vk *VerifyingKey) writeTo(w io.Writer, options ...func(*curve.Encoder)) (n int64, err error) {
	enc := curve.NewEncoder(w)

	toEncode := []interface{}{
		vk.Size,
		&vk.SizeInv,
		&vk.Generator,
		vk.NbPublicVariables,
		&vk.CosetShift,
		&vk.Qpublic,
		&vk.Kzg.G1,
		&vk.Kzg.G2[0],
		&vk.Kzg.G2[1],
		&vk.Kzg.Lines,
		vk.CommitmentConstraintIndexes,
	}

	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			return enc.BytesWritten(), err
		}
	}

	return enc.BytesWritten(), nil
}

// UnsafeReadFrom reads from binary representation in r into VerifyingKey.
// Current implementation is a passthrough to ReadFrom
func (vk *VerifyingKey) UnsafeReadFrom(r io.Reader) (int64, error) {
	return vk.ReadFrom(r)
}

// ReadFrom reads from binary representation in r into VerifyingKey
func (vk *VerifyingKey) ReadFrom(r io.Reader) (int64, error) {
	dec := curve.NewDecoder(r)
	toDecode := []interface{}{
		&vk.Size,
		&vk.SizeInv,
		&vk.Generator,
		&vk.NbPublicVariables,
		&vk.CosetShift,
		&vk.Qpublic,
		&vk.Kzg.G1,
		&vk.Kzg.G2[0],
		&vk.Kzg.G2[1],
		&vk.Kzg.Lines,
		&vk.CommitmentConstraintIndexes,
	}

	for _, v := range toDecode {
		if err := dec.Decode(v); err != nil {
			return dec.BytesRead(), err
		}
	}

	return dec.BytesRead(), nil
}
