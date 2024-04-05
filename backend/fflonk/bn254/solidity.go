package fflonk

const tmplSolidityVerifier = `// SPDX-License-Identifier: Apache-2.0

// Copyright 2023 Consensys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gnark DO NOT EDIT

pragma solidity ^0.8.19;

contract PlonkVerifier {

  uint256 private constant R_MOD = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
  uint256 private constant R_MOD_MINUS_ONE = 21888242871839275222246405745257275088548364400416034343698204186575808495616;
  uint256 private constant P_MOD = 21888242871839275222246405745257275088696311157297823662689037894645226208583;
  {{- range $index, $element := .Kzg.G2 }}
  uint256 private constant G2_SRS_{{ $index }}_X_0 = {{ (fpstr $element.X.A1) }};
  uint256 private constant G2_SRS_{{ $index }}_X_1 = {{ (fpstr $element.X.A0) }};
  uint256 private constant G2_SRS_{{ $index }}_Y_0 = {{ (fpstr $element.Y.A1) }};
  uint256 private constant G2_SRS_{{ $index }}_Y_1 = {{ (fpstr $element.Y.A0) }};
  {{ end -}}
  uint256 private constant G1_SRS_X = {{ fpstr .Kzg.G1.X }};
  uint256 private constant G1_SRS_Y = {{ fpstr .Kzg.G1.Y }};

  // ----------------------- vk ---------------------
  uint256 private constant VK_DOMAIN_SIZE = {{ .Size }};
  uint256 private constant VK_INV_DOMAIN_SIZE = {{ (frstr .SizeInv) }};
  uint256 private constant VK_OMEGA = {{ (frstr .Generator) }};
  uint256 private constant VK_NB_PUBLIC_INPUTS = {{ .NbPublicVariables }};
  uint256 private constant VK_COSET_SHIFT = 5;
  uint256 private constant VK_QPUBLIC_COM_X = {{ (fpstr .Qpublic.X) }};
  uint256 private constant VK_QPUBLIC_COM_Y = {{ (fpstr .Qpublic.Y) }};
  {{ range $index, $element := .CommitmentConstraintIndexes -}}
  uint256 private constant VK_INDEX_COMMIT_API_{{ $index }} = {{ $element }};
  {{ end -}}
  uint256 private constant VK_NB_CUSTOM_GATES = {{ len .CommitmentConstraintIndexes }};
  uint256 private constant T_TH_ROOT_ONE = {{ tThRootOne . }};
  uint256 private constant VK_T = {{ nextDivisorRMinusOne . }};

  // --------------------------- proof -----------------

  uint256 private constant PROOF_LROENTANGLED_COM_X = 0x00;
  uint256 private constant PROOF_LROENTANGLED_COM_Y = 0x20;
  uint256 private constant PROOF_Z_X = 0x40;
  uint256 private constant PROOF_Z_Y = 0x60;
  uint256 private constant PROOF_Z_ENTANGLED_X = 0x80;
  uint256 private constant PROOF_Z_ENTANGLED_Y = 0xa0;
  uint256 private constant PROOF_H_ENTANGLED_X = 0xc0;
  uint256 private constant PROOF_H_ENTANGLED_Y = 0xe0;
  {{- range $index, $element := .CommitmentConstraintIndexes}}
  uint256 private constant PROOF_BSB_{{ $index }}_X = {{ hex ( add 0x100 (mul 0x20 $index) ) }};
  uint256 private constant PROOF_BSB_{{ $index }}_Y = {{ hex ( add 0x120 (mul 0x20 $index) ) }};
  {{ end -}}{{ $offset := add 0x100 (mul 0x40 (len .CommitmentConstraintIndexes )) }}
  uint256 private constant PROOF_QL_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_QR_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_QM_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_QO_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_QKINCOMPLETE_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_S1_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_S2_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_S3_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  {{- range $index, $element := .CommitmentConstraintIndexes}}
  uint256 private constant PROOF_QCP_{{ $index }}_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  {{ end -}}
  uint256 private constant PROOF_L_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_R_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_O_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_Z_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_H1_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_H2_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_H3_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  {{- range $index, $element := .CommitmentConstraintIndexes}}
  uint256 private constant PROOF_BSB_{{ $index }}_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  {{ end -}}
  uint256 private constant PROOF_Z_AT_ZETA_T_OMEGA = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_W_X = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_W_Y = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_W_PRIME_X = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_W_PRIME_Y = {{ hex $offset }};{{ $offset = add $offset 0x20}}

  uint256 private constant PROOF_SHPLONK_P0_0 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_1 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_2 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_3 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_4 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_5 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_6 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_7 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_8 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_9 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_10 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_11 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_12 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_13 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_14 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_15 = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  {{- range $index, $element := .CommitmentConstraintIndexes }}
  uint256 private constant PROOF_SHPLONK_P0_{{ add 16 (mul 2 $index)}} = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant PROOF_SHPLONK_P0_{{ add 17 (mul 2 $index)}} = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  {{ end -}}
  uint256 private constant PROOF_SHPLONK_P1_0 = {{ hex $offset }};{{ $offset = add $offset 0x20}}

  // -------- offset state {{ $offset = 0 }}
  uint256 private constant STATE_ALPHA = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant STATE_BETA = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant STATE_GAMMA = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant STATE_ZETA = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant STATE_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant STATE_ZH_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant STATE_LAGRANGE_0_AT_ZETA_T = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant STATE_PI = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant STATE_SUCCESS = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant STATE_CHECK_VAR = {{ hex $offset }};{{ $offset = add $offset 0x20}}
  uint256 private constant STATE_LAST_MEM = {{ hex $offset }};{{ $offset = add $offset 0x20}}

  // -------- errors
  uint256 private constant ERROR_STRING_ID = 0x08c379a000000000000000000000000000000000000000000000000000000000; // selector for function Error(string)

  {{ if (gt (len .CommitmentConstraintIndexes) 0 )}}
  // -------- utils (for hash_fr)
	uint256 private constant HASH_FR_BB = 340282366920938463463374607431768211456; // 2**128
	uint256 private constant HASH_FR_ZERO_UINT256 = 0;
	uint8 private constant HASH_FR_LEN_IN_BYTES = 48;
	uint8 private constant HASH_FR_SIZE_DOMAIN = 11;
	uint8 private constant HASH_FR_ONE = 1;
	uint8 private constant HASH_FR_TWO = 2;
  {{ end }}

  // -------- precompiles
  uint256 private constant SHA_256 = 0x2;
  uint256 private constant MOD_EXP = 0x5;
  uint256 private constant EC_ADD = 0x6;
  uint256 private constant EC_MUL = 0x7;
  uint256 private constant EC_PAIRING = 0x8;

  /// Verify a Plonk proof.
  /// Reverts if the proof or the public inputs are malformed.
  /// @param proof serialised plonk proof (using gnark's MarshalSolidity)
  /// @param public_inputs (must be reduced)
  /// @return success true if the proof passes false otherwise
  function Verify(bytes calldata proof, uint256[] calldata public_inputs) 
  public view returns(bool success) {

	assembly {

		// state memory and scratch memory
		let mem := mload(0x40)
		let freeMem := add(mem, STATE_LAST_MEM)

		// compute the challenges
		let prev_challenge_non_reduced
		prev_challenge_non_reduced := derive_gamma(proof.offset, public_inputs.length, public_inputs.offset)
		prev_challenge_non_reduced := derive_beta(prev_challenge_non_reduced)
		prev_challenge_non_reduced := derive_alpha(proof.offset, prev_challenge_non_reduced)
		derive_zeta(proof.offset, prev_challenge_non_reduced)
    compute_zh_zeta_t(freeMem)

    let l_pi := compute_pi(public_inputs.offset, public_inputs.length, freeMem)
    {{ if (gt (len .CommitmentConstraintIndexes) 0 ) -}}
    let l_pi_with_commit := compute_pi_commit(proof.offset, public_inputs.length, freeMem)
    l_pi := addmod(l_pi_with_commit, l_pi, R_MOD)
    {{ end -}}
    mstore(add(mem, STATE_PI), l_pi)

		// Beginning errors -------------------------------------------------

    function error_nb_public_inputs() {
      let ptError := mload(0x40)
      mstore(ptError, ERROR_STRING_ID) // selector for function Error(string)
      mstore(add(ptError, 0x4), 0x20)
      mstore(add(ptError, 0x24), 0x1d)
      mstore(add(ptError, 0x44), "wrong number of public inputs")
      revert(ptError, 0x64)
    }

    /// Called when an operation on Bn254 fails
    /// @dev for instance when calling EcMul on a point not on Bn254.
    function error_ec_op() {
      let ptError := mload(0x40)
      mstore(ptError, ERROR_STRING_ID) // selector for function Error(string)
      mstore(add(ptError, 0x4), 0x20)
      mstore(add(ptError, 0x24), 0x12)
      mstore(add(ptError, 0x44), "error ec operation")
      revert(ptError, 0x64)
    }

    /// Called when one of the public inputs is not reduced.
    function error_inputs_size() {
      let ptError := mload(0x40)
      mstore(ptError, ERROR_STRING_ID) // selector for function Error(string)
      mstore(add(ptError, 0x4), 0x20)
      mstore(add(ptError, 0x24), 0x18)
      mstore(add(ptError, 0x44), "inputs are bigger than r")
      revert(ptError, 0x64)
    }

    /// Called when the size proof is not as expected
    /// @dev to avoid overflow attack for instance
    function error_proof_size() {
      let ptError := mload(0x40)
      mstore(ptError, ERROR_STRING_ID) // selector for function Error(string)
      mstore(add(ptError, 0x4), 0x20)
      mstore(add(ptError, 0x24), 0x10)
      mstore(add(ptError, 0x44), "wrong proof size")
      revert(ptError, 0x64)
    }

    /// Called when one the openings is bigger than r
    /// The openings are the claimed evalutions of a polynomial
    /// in a Kzg proof.
    function error_proof_openings_size() {
      let ptError := mload(0x40)
      mstore(ptError, ERROR_STRING_ID) // selector for function Error(string)
      mstore(add(ptError, 0x4), 0x20)
      mstore(add(ptError, 0x24), 0x16)
      mstore(add(ptError, 0x44), "openings bigger than r")
      revert(ptError, 0x64)
    }

    function error_verify() {
      let ptError := mload(0x40)
      mstore(ptError, ERROR_STRING_ID) // selector for function Error(string)
      mstore(add(ptError, 0x4), 0x20)
      mstore(add(ptError, 0x24), 0xc)
      mstore(add(ptError, 0x44), "error verify")
      revert(ptError, 0x64)
    }

    function error_random_generation() {
      let ptError := mload(0x40)
      mstore(ptError, ERROR_STRING_ID) // selector for function Error(string)
      mstore(add(ptError, 0x4), 0x20)
      mstore(add(ptError, 0x24), 0x14)
      mstore(add(ptError, 0x44), "error random gen kzg")
      revert(ptError, 0x64)
    }
    // end errors -------------------------------------------------

		// Beginning challenges -------------------------------------------------

		/// Derive gamma as Sha256(<transcript>)
		/// @param aproof pointer to the proof
		/// @param nb_pi number of public inputs
		/// @param pi pointer to the array of public inputs
		/// @return the challenge gamma, not reduced
		/// @notice The transcript is the concatenation (in this order) of:
		/// * the word "gamma" in ascii, equal to [0x67,0x61,0x6d, 0x6d, 0x61] and encoded as a uint256.
		/// * the entangled commitments to ql,qr,qm,qo,qkIncomplete,s1,s2,s3,qcp_i
		/// * the entangled commitment to l, r, o
		/// The data described above is written starting at mPtr. "gamma" lies on 5 bytes,
		/// and is encoded as a uint256 number n. In basis b = 256, the number looks like this
		/// [0 0 0 .. 0x67 0x61 0x6d, 0x6d, 0x61]. The first non zero entry is at position 27=0x1b
		/// Gamma reduced (the actual challenge) is stored at add(state, state_gamma)
		function derive_gamma(aproof, nb_pi, pi)->gamma_not_reduced {
			let state := mload(0x40)
			let mPtr := add(state, STATE_LAST_MEM)
			mstore(mPtr, 0x67616d6d61) // "gamma" in ascii is [0x67,0x61,0x6d, 0x6d, 0x61]
			mstore(add(mPtr, 0x20), VK_QPUBLIC_COM_X)
			mstore(add(mPtr, 0x40), VK_QPUBLIC_COM_Y)
			let _mPtr := add(mPtr, 0x60)
			let size_pi_in_bytes := mul(nb_pi, 0x20)
			calldatacopy(_mPtr, pi, size_pi_in_bytes)
			_mPtr := add(_mPtr, size_pi_in_bytes)
			calldatacopy(_mPtr, add(aproof, PROOF_LROENTANGLED_COM_X), 0x40)
			let size := add(0x85, size_pi_in_bytes)
			
			let l_success := staticcall(gas(), SHA_256, add(mPtr, 0x1b), size, mPtr, 0x20) //0x1b -> 000.."gamma"
			if iszero(l_success) {
			error_verify()
			}
			gamma_not_reduced := mload(mPtr)
			mstore(add(state, STATE_GAMMA), mod(gamma_not_reduced, R_MOD))
		}

		/// derive beta as Sha256<transcript>
		/// @param gamma_not_reduced the previous challenge (gamma) not reduced
		/// @return beta_not_reduced the next challenge, beta, not reduced
		/// @notice the transcript consists of the previous challenge only.
		/// The reduced version of beta is stored at add(state, state_beta)
		function derive_beta(gamma_not_reduced)->beta_not_reduced{
      let state := mload(0x40)
			let mPtr := add(mload(0x40), STATE_LAST_MEM)
			mstore(mPtr, 0x62657461) // "beta" in ascii is [0x62,0x65,0x74,0x61]
			mstore(add(mPtr, 0x20), gamma_not_reduced)
			let l_success := staticcall(gas(), SHA_256, add(mPtr, 0x1c), 0x24, mPtr, 0x20) //0x1b -> 000.."gamma"
			if iszero(l_success) {
			error_verify()
			}
			beta_not_reduced := mload(mPtr)
			mstore(add(state, STATE_BETA), mod(beta_not_reduced, R_MOD))
		}

		/// derive alpha as sha256<transcript>
		/// @param aproof pointer to the proof object
		/// @param beta_not_reduced the previous challenge (beta) not reduced
		/// @return alpha_not_reduced the next challenge, alpha, not reduced
		/// @notice the transcript consists of the previous challenge (beta)
		/// not reduced, the commitments to the wires associated to the QCP_i,
		/// and the commitment to the grand product polynomial 
		function derive_alpha(aproof, beta_not_reduced)->alpha_not_reduced {
			let state := mload(0x40)
			let mPtr := add(mload(0x40), STATE_LAST_MEM)
			mstore(mPtr, 0x616C706861) // "alpha" in ascii is [0x61,0x6C,0x70,0x68,0x61]
			let _mPtr := add(mPtr, 0x20)
			mstore(_mPtr, beta_not_reduced)
			_mPtr := add(_mPtr, 0x20)
			let size_bsb_commitments := mul(0x40, VK_NB_CUSTOM_GATES)
			calldatacopy(_mPtr, add(aproof, PROOF_BSB_0_X), size_bsb_commitments)
			_mPtr := add(_mPtr, size_bsb_commitments)
			calldatacopy(_mPtr, add(aproof, PROOF_Z_ENTANGLED_X), 0x40)
      let size := add(0x65, size_bsb_commitments)
			let l_success := staticcall(gas(), SHA_256, add(mPtr, 0x1b), size, mPtr, 0x20)
			if iszero(l_success) {
			error_verify()
			}

			alpha_not_reduced := mload(mPtr)
			mstore(add(state, STATE_ALPHA), mod(alpha_not_reduced, R_MOD))
		}

		/// derive zeta as sha256<transcript>
		/// @param aproof pointer to the proof object
		/// @param alpha_not_reduced the previous challenge (alpha) not reduced
		/// The transcript consists of the previous challenge and the commitment to
		/// the quotient polynomial h.
		function derive_zeta(aproof, alpha_not_reduced) {
			let state := mload(0x40)
			let mPtr := add(mload(0x40), STATE_LAST_MEM)
			mstore(mPtr, 0x7a657461) // "zeta" in ascii is [0x7a,0x65,0x74,0x61]
			mstore(add(mPtr, 0x20), alpha_not_reduced)
			calldatacopy(add(mPtr, 0x40), add(aproof, PROOF_H_ENTANGLED_X), 0xc0)
			let l_success := staticcall(gas(), SHA_256, add(mPtr, 0x1c), 0x64, mPtr, 0x20)
			if iszero(l_success) {
			error_verify()
			}
			let zeta_not_reduced := mload(mPtr)
			mstore(add(state, STATE_ZETA), mod(zeta_not_reduced, R_MOD))
		}
		// END challenges -------------------------------------------------

    // BEGINNING compute_pi -------------------------------------------------

    /// compute ζᵗ and
    /// computes  Z=Xⁿ-1 at ζᵗ
    function compute_zh_zeta_t(mPtr) {
      let state := mload(0x40)
      let zeta := mload(add(state, STATE_ZETA))
      let zh_zeta_t := pow(zeta, VK_T, mPtr)
      mstore(add(state, STATE_ZETA_T), zh_zeta_t)
      zh_zeta_t := addmod(pow(zh_zeta_t, VK_DOMAIN_SIZE, mPtr), R_MOD_MINUS_ONE, R_MOD)
      mstore(add(state, STATE_ZH_ZETA_T), zh_zeta_t)
    }


    /// compute_pi computes the public inputs contributions,
    /// except for the public inputs coming from the custom gate
    /// @param ins pointer to the public inputs
    /// @param n number of public inputs
    /// @param mPtr free memory
    /// @return pi_wo_commit public inputs contribution (except the public inputs coming from the custom gate)
    function compute_pi(ins, n, mPtr)-check_constraints>pi_wo_commit {
      let state := mload(0x40)
      let zt := mload(add(state, STATE_ZETA_T))
      let zh_zeta_t := mload(add(state, STATE_ZH_ZETA_T))
      let li := mPtr
      batch_compute_lagranges_at_z(zt, zh_zeta_t, n, li)
      mstore(add(state, STATE_LAGRANGE_0_AT_ZETA_T), mload(li))
      let tmp := 0
      for {let i:=0} lt(i,n) {i:=add(i,1)}
      {
        tmp := mulmod(mload(li), calldataload(ins), R_MOD)
        pi_wo_commit := addmod(pi_wo_commit, tmp, R_MOD)
        li := add(li, 0x20)
        ins := add(ins, 0x20)
      } 
    }

    /// batch_compute_lagranges_at_z computes [L_0(z), .., L_{n-1}(z)]
    /// @param z point at which the Lagranges are evaluated
    /// @param zh_zeta_t ζⁿ-1
    /// @param n number of public inputs (number of Lagranges to compute)
    /// @param mPtr pointer to which the results are stored
    function batch_compute_lagranges_at_z(z, zh_zeta_t, n, mPtr) {
      let zn := mulmod(zh_zeta_t, VK_INV_DOMAIN_SIZE, R_MOD) // 1/n * (ζⁿ - 1)
      let _w := 1
      let _mPtr := mPtr
      for {let i:=0} lt(i,n) {i:=add(i,1)}
      {
        mstore(_mPtr, addmod(z,sub(R_MOD, _w), R_MOD))
        _w := mulmod(_w, VK_OMEGA, R_MOD)
        _mPtr := add(_mPtr, 0x20)
      }
      batch_invert(mPtr, n, _mPtr)
      _mPtr := mPtr
      _w := 1
      for {let i:=0} lt(i,n) {i:=add(i,1)}
      {
        mstore(_mPtr, mulmod(mulmod(mload(_mPtr), zn , R_MOD), _w, R_MOD))
        _mPtr := add(_mPtr, 0x20)
        _w := mulmod(_w, VK_OMEGA, R_MOD)
      }
    } 

    /// @notice Montgomery trick for batch inversion mod R_MOD
    /// @param ins pointer to the data to batch invert
    /// @param number of elements to batch invert
    /// @param mPtr free memory
    function batch_invert(ins, nb_ins, mPtr) {
      mstore(mPtr, 1)
      let offset := 0
      for {let i:=0} lt(i, nb_ins) {i:=add(i,1)}
      {
        let prev := mload(add(mPtr, offset))
        let cur := mload(add(ins, offset))
        cur := mulmod(prev, cur, R_MOD)
        offset := add(offset, 0x20)
        mstore(add(mPtr, offset), cur)
      }
      ins := add(ins, sub(offset, 0x20))
      mPtr := add(mPtr, offset)
      let inv := pow(mload(mPtr), sub(R_MOD,2), add(mPtr, 0x20))
      for {let i:=0} lt(i, nb_ins) {i:=add(i,1)}
      {
        mPtr := sub(mPtr, 0x20)
        let tmp := mload(ins)
        let cur := mulmod(inv, mload(mPtr), R_MOD)
        mstore(ins, cur)
        inv := mulmod(inv, tmp, R_MOD)
        ins := sub(ins, 0x20)
      }
    }

    {{ if (gt (len .CommitmentConstraintIndexes) 0 )}}
    /// Public inputs (the ones coming from the custom gate) contribution
    /// @param aproof pointer to the proof
    /// @param nb_public_inputs number of public inputs
    /// @param mPtr pointer to free memory
    /// @return pi_commit custom gate public inputs contribution
    function compute_pi_commit(aproof, nb_public_inputs, mPtr)->pi_commit {
      let state := mload(0x40)
      let zt := mload(add(state, STATE_ZETA_T))
      let zh_zeta_t := mload(add(state, STATE_ZH_ZETA_T))
      let p := add(aproof, PROOF_BSB_0_X)
      let h_fr, ith_lagrange
      {{ range $index, $element := .CommitmentConstraintIndexes}}
      h_fr := hash_fr(calldataload(p), calldataload(add(p, 0x20)), mPtr)
      ith_lagrange := compute_ith_lagrange_at_z(zt, zh_zeta_t, add(nb_public_inputs, VK_INDEX_COMMIT_API_{{ $index }}), mPtr)
      pi_commit := addmod(pi_commit, mulmod(h_fr, ith_lagrange, R_MOD), R_MOD)
      p := add(p, 0x40)
      {{ end }}
    }

    /// Computes L_i(zeta) =  ωⁱ/n * (ζᵗⁿ-1)/(ζᵗ-ωⁱ) where:
    /// @param ζᵗ zeta
    /// @param zt ζᵗⁿ-1
    /// @param i i-th lagrange
    /// @param mPtr free memory
    /// @return res = ωⁱ/n * (ζᵗⁿ-1)/(ζᵗ-ωⁱ) 
    function compute_ith_lagrange_at_z(zt, zh_zeta_t, i, mPtr)->res {
      let w := pow(VK_OMEGA, i, mPtr) // w**i
      i := addmod(zt, sub(R_MOD, w), R_MOD) // z-w**i
      w := mulmod(w, VK_INV_DOMAIN_SIZE, R_MOD) // w**i/n
      i := pow(i, sub(R_MOD,2), mPtr) // (z-w**i)**-1
      w := mulmod(w, i, R_MOD) // w**i/n*(z-w)**-1
      res := mulmod(w, zh_zeta_t, R_MOD)
    }

    /// @dev https://tools.ietf.org/html/draft-irtf-cfrg-hash-to-curve-06#section-5.2
    /// @param x x coordinate of a point on Bn254(𝔽_p)
    /// @param y y coordinate of a point on Bn254(𝔽_p)
    /// @param mPtr free memory
    /// @return res an element mod R_MOD
    function hash_fr(x, y, mPtr)->res {

      // [0x00, .. , 0x00 || x, y, || 0, 48, 0, dst, HASH_FR_SIZE_DOMAIN]
      // <-  64 bytes  ->  <-64b -> <-       1 bytes each     ->

      // [0x00, .., 0x00] 64 bytes of zero
      mstore(mPtr, HASH_FR_ZERO_UINT256)
      mstore(add(mPtr, 0x20), HASH_FR_ZERO_UINT256)
  
      // msg =  x || y , both on 32 bytes
      mstore(add(mPtr, 0x40), x)
      mstore(add(mPtr, 0x60), y)

      // 0 || 48 || 0 all on 1 byte
      mstore8(add(mPtr, 0x80), 0)
      mstore8(add(mPtr, 0x81), HASH_FR_LEN_IN_BYTES)
      mstore8(add(mPtr, 0x82), 0)

      // "BSB22-Plonk" = [42, 53, 42, 32, 32, 2d, 50, 6c, 6f, 6e, 6b,]
      mstore8(add(mPtr, 0x83), 0x42)
      mstore8(add(mPtr, 0x84), 0x53)
      mstore8(add(mPtr, 0x85), 0x42)
      mstore8(add(mPtr, 0x86), 0x32)
      mstore8(add(mPtr, 0x87), 0x32)
      mstore8(add(mPtr, 0x88), 0x2d)
      mstore8(add(mPtr, 0x89), 0x50)
      mstore8(add(mPtr, 0x8a), 0x6c)
      mstore8(add(mPtr, 0x8b), 0x6f)
      mstore8(add(mPtr, 0x8c), 0x6e)
      mstore8(add(mPtr, 0x8d), 0x6b)

      // size domain
      mstore8(add(mPtr, 0x8e), HASH_FR_SIZE_DOMAIN)

      let l_success := staticcall(gas(), 0x2, mPtr, 0x8f, mPtr, 0x20)
      if iszero(l_success) {
        error_verify()
      }

      let b0 := mload(mPtr)

      // [b0         || one || dst || HASH_FR_SIZE_DOMAIN]
      // <-64bytes ->  <-    1 byte each      ->
      mstore8(add(mPtr, 0x20), HASH_FR_ONE) // 1
      
      mstore8(add(mPtr, 0x21), 0x42) // dst
      mstore8(add(mPtr, 0x22), 0x53)
      mstore8(add(mPtr, 0x23), 0x42)
      mstore8(add(mPtr, 0x24), 0x32)
      mstore8(add(mPtr, 0x25), 0x32)
      mstore8(add(mPtr, 0x26), 0x2d)
      mstore8(add(mPtr, 0x27), 0x50)
      mstore8(add(mPtr, 0x28), 0x6c)
      mstore8(add(mPtr, 0x29), 0x6f)
      mstore8(add(mPtr, 0x2a), 0x6e)
      mstore8(add(mPtr, 0x2b), 0x6b)

      mstore8(add(mPtr, 0x2c), HASH_FR_SIZE_DOMAIN) // size domain
      l_success := staticcall(gas(), 0x2, mPtr, 0x2d, mPtr, 0x20)
      if iszero(l_success) {
        error_verify()
      }

      // b1 is located at mPtr. We store b2 at add(mPtr, 0x20)

      // [b0^b1      || two || dst || HASH_FR_SIZE_DOMAIN]
      // <-64bytes ->  <-    1 byte each      ->
      mstore(add(mPtr, 0x20), xor(mload(mPtr), b0))
      mstore8(add(mPtr, 0x40), HASH_FR_TWO)

      mstore8(add(mPtr, 0x41), 0x42) // dst
      mstore8(add(mPtr, 0x42), 0x53)
      mstore8(add(mPtr, 0x43), 0x42)
      mstore8(add(mPtr, 0x44), 0x32)
      mstore8(add(mPtr, 0x45), 0x32)
      mstore8(add(mPtr, 0x46), 0x2d)
      mstore8(add(mPtr, 0x47), 0x50)
      mstore8(add(mPtr, 0x48), 0x6c)
      mstore8(add(mPtr, 0x49), 0x6f)
      mstore8(add(mPtr, 0x4a), 0x6e)
      mstore8(add(mPtr, 0x4b), 0x6b)

      mstore8(add(mPtr, 0x4c), HASH_FR_SIZE_DOMAIN) // size domain

      let offset := add(mPtr, 0x20)
      l_success := staticcall(gas(), 0x2, offset, 0x2d, offset, 0x20)
      if iszero(l_success) {
        error_verify()
      }

      // at this point we have mPtr = [ b1 || b2] where b1 is on 32byes and b2 in 16bytes.
      // we interpret it as a big integer mod r in big endian (similar to regular decimal notation)
      // the result is then 2**(8*16)*mPtr[32:] + mPtr[32:48]
      res := mulmod(mload(mPtr), HASH_FR_BB, R_MOD) // <- res = 2**128 * mPtr[:32]
      let b1 := shr(128, mload(add(mPtr, 0x20))) // b1 <- [0, 0, .., 0 ||  b2[:16] ]
      res := addmod(res, b1, R_MOD)

    }
    {{ end }}
    // END compute_pi -------------------------------------------------

    // BEGINNING utils math functions -------------------------------------------------
      
    /// @param dst pointer storing the result
    /// @param p pointer to the first point
    /// @param q pointer to the second point
    /// @param mPtr pointer to free memory
    function point_add(dst, p, q, mPtr) {
      let state := mload(0x40)
      mstore(mPtr, mload(p))
      mstore(add(mPtr, 0x20), mload(add(p, 0x20)))
      mstore(add(mPtr, 0x40), mload(q))
      mstore(add(mPtr, 0x60), mload(add(q, 0x20)))
      let l_success := staticcall(gas(),EC_ADD,mPtr,0x80,dst,0x40)
      if iszero(l_success) {
        error_ec_op()
      }
    }

    /// @param dst pointer storing the result
    /// @param p pointer to the first point (calldata)
    /// @param q pointer to the second point (calladata)
    /// @param mPtr pointer to free memory
    function point_add_calldata(dst, p, q, mPtr) {
      let state := mload(0x40)
      mstore(mPtr, mload(p))
      mstore(add(mPtr, 0x20), mload(add(p, 0x20)))
      mstore(add(mPtr, 0x40), calldataload(q))
      mstore(add(mPtr, 0x60), calldataload(add(q, 0x20)))
      let l_success := staticcall(gas(), EC_ADD, mPtr, 0x80, dst, 0x40)
      if iszero(l_success) {
        error_ec_op()
      }
    }

    /// @parma dst pointer storing the result
    /// @param src pointer to a point on Bn254(𝔽_p)
    /// @param s scalar
    /// @param mPtr free memory
    function point_mul(dst,src,s, mPtr) {
      let state := mload(0x40)
      mstore(mPtr,mload(src))
      mstore(add(mPtr,0x20),mload(add(src,0x20)))
      mstore(add(mPtr,0x40),s)
      let l_success := staticcall(gas(),EC_MUL,mPtr,0x60,dst,0x40)
      if iszero(l_success) {
        error_ec_op()
      }
    }

    /// @parma dst pointer storing the result
    /// @param src pointer to a point on Bn254(𝔽_p) on calldata
    /// @param s scalar
    /// @param mPtr free memory
    function point_mul_calldata(dst, src, s, mPtr) {
      let state := mload(0x40)
      mstore(mPtr, calldataload(src))
      mstore(add(mPtr, 0x20), calldataload(add(src, 0x20)))
      mstore(add(mPtr, 0x40), s)
      let l_success := staticcall(gas(), EC_MUL, mPtr, 0x60, dst, 0x40)
      if iszero(l_success) {
        error_ec_op()
      }
    }

    /// @notice dst <- dst + [s]src (Elliptic curve)
    /// @param dst pointer accumulator point storing the result
    /// @param src pointer to the point to multiply and add
    /// @param s scalar
    /// @param mPtr free memory
    function point_acc_mul(dst,src,s, mPtr) {
      let state := mload(0x40)
      mstore(mPtr,mload(src))
      mstore(add(mPtr,0x20),mload(add(src,0x20)))
      mstore(add(mPtr,0x40),s)
      let l_success := staticcall(gas(),EC_MUL,mPtr,0x60,mPtr,0x40)
      mstore(add(mPtr,0x40),mload(dst))
      mstore(add(mPtr,0x60),mload(add(dst,0x20)))
      l_success := and(l_success, staticcall(gas(),EC_ADD,mPtr,0x80,dst, 0x40))
      if iszero(l_success) {
        error_ec_op()
      }
    }

    /// @notice dst <- dst + [s]src (Elliptic curve)
    /// @param dst pointer accumulator point storing the result
    /// @param src pointer to the point to multiply and add (on calldata)
    /// @param s scalar
    /// @mPtr free memory
    function point_acc_mul_calldata(dst, src, s, mPtr) {
      let state := mload(0x40)
      mstore(mPtr, calldataload(src))
      mstore(add(mPtr, 0x20), calldataload(add(src, 0x20)))
      mstore(add(mPtr, 0x40), s)
      let l_success := staticcall(gas(), EC_MUL, mPtr, 0x60, mPtr, 0x40)
      mstore(add(mPtr, 0x40), mload(dst))
      mstore(add(mPtr, 0x60), mload(add(dst, 0x20)))
      l_success := and(l_success, staticcall(gas(), EC_ADD, mPtr, 0x80, dst, 0x40))
      if iszero(l_success) {
        error_ec_op()
      }
    }

    /// @notice dst <- dst + src*s (Fr) dst,src are addresses, s is a value
    /// @param dst pointer storing the result
    /// @param src pointer to the scalar to multiply and add (on calldata)
    /// @param s scalar
    function fr_acc_mul_calldata(dst, src, s) {
      let tmp :=  mulmod(calldataload(src), s, R_MOD)
      mstore(dst, addmod(mload(dst), tmp, R_MOD))
    }

    /// @param x element to exponentiate
    /// @param e exponent
    /// @param mPtr free memory
    /// @return res x ** e mod r
    function pow(x, e, mPtr)->res {
      mstore(mPtr, 0x20)
      mstore(add(mPtr, 0x20), 0x20)
      mstore(add(mPtr, 0x40), 0x20)
      mstore(add(mPtr, 0x60), x)
      mstore(add(mPtr, 0x80), e)
      mstore(add(mPtr, 0xa0), R_MOD)
      let check_staticcall := staticcall(gas(),MOD_EXP,mPtr,0xc0,mPtr,0x20)
      if eq(check_staticcall, 0) {
        error_verify()
      }
      res := mload(mPtr)
    }
    // end math functions -------------------------------------------------

    function check_constraints(aproof) {
      let state := mload(0x40)
      let gates := check_gates(aproof)
      let permutation := check_permutation(aproof)
      let start_at_one := check_z_start_at_one(aproof)
      let alpha := mload(add(state, STATE_ALPHA))
      let lhs := mulmod(start_at_one, alpha, R_MOD)
      lhs := addmod(lhs, permutation, R_MOD)
      lhs := mulmod(lhs, alpha, R_MOD)
      lhs := addmod(lhs, gates, R_MOD)
      
      let mPtr := mload(add(state, STATE_LAST_MEM))
      let zt := mload(add(state, STATE_ZETA_T))
      let zh_zeta_t := mload(add(state, STATE_ZH_ZETA_T))
      let zzt := addmod(zh_zeta_t, 1, R_MOD)
      zzt := mulmod(zt, mulmod(zzt, zt, R_MOD), R_MOD)
      let rhs := calldataload(add(aproof, PROOF_H3_AT_ZETA_T))
      rhs := mulmod(rhs, zzt, R_MOD)
      rhs := addmod(rhs, calldataload(add(aproof, PROOF_H2_AT_ZETA_T)), R_MOD)
      rhs := mulmod(rhs, zzt, R_MOD)
      rhs := addmod(rhs, calldataload(add(aproof, PROOF_H1_AT_ZETA_T)), R_MOD)
      rhs := mulmod(rhs, zh_zeta_t, R_MOD)

      if iszero(eq(lhs, rhs)) {
        error_verify()
      }
    }

    function check_gates(aproof)->ag {
      let state := mload(0x40)
      let ql := calldataload(add(aproof, PROOF_QL_AT_ZETA_T))
      let qr := calldataload(add(aproof, PROOF_QR_AT_ZETA_T))
      let qm := calldataload(add(aproof, PROOF_QM_AT_ZETA_T))
      let qo := calldataload(add(aproof, PROOF_QO_AT_ZETA_T))
      let qk := calldataload(add(aproof, PROOF_QKINCOMPLETE_AT_ZETA_T))
      let l := calldataload(add(aproof, PROOF_L_AT_ZETA_T))
      let r := calldataload(add(aproof, PROOF_R_AT_ZETA_T))
      let o := calldataload(add(aproof, PROOF_O_AT_ZETA_T))
      ag := mulmod(ql, l, R_MOD)
      let tmp := mulmod(qr, r, R_MOD)
      ag := addmod(ag, tmp, R_MOD)
      tmp := mulmod(mulmod(qm, l, R_MOD), r, R_MOD)
      ag := addmod(ag, tmp, R_MOD)
      tmp := mulmod(qo, o, R_MOD)
      ag := addmod(ag, tmp, R_MOD)
      tmp := addmod(qk, mload(add(state, STATE_PI)), R_MOD)
      ag := addmod(ag, tmp, R_MOD)
      {{ range $index, $element := .CommitmentConstraintIndexes -}}
      tmp := mulmod(calldataload(add(aproof, PROOF_QCP_{{$index}}_AT_ZETA_T)), calldataload(add(aproof, PROOF_BSB_{{$index}}_AT_ZETA_T)), R_MOD)
      ag := addmod(ag, tmp, R_MOD)
      {{ end -}}
    }

    function check_permutation(aproof)->permutation {
      let state := mload(0x40)
      let zeta_t := mload(add(state, STATE_ZETA_T))
      let u_zeta_t := mulmod(zeta_t, VK_COSET_SHIFT, R_MOD)
      let uu_zeta_t := mulmod(u_zeta_t, VK_COSET_SHIFT, R_MOD)
      let beta := mload(add(state, STATE_BETA))
      let gamma := mload(add(state, STATE_GAMMA))
      let l := calldataload(add(aproof, PROOF_L_AT_ZETA_T))
      let r := calldataload(add(aproof, PROOF_R_AT_ZETA_T))
      let o := calldataload(add(aproof, PROOF_O_AT_ZETA_T))

      permutation := mulmod(beta, calldataload(add(aproof, PROOF_S1_AT_ZETA_T)), R_MOD)
      permutation := addmod(permutation, l, R_MOD)
      permutation := addmod(permutation, gamma, R_MOD)
      let tmp := mulmod(beta, calldataload(add(aproof, PROOF_S2_AT_ZETA_T)), R_MOD)
      tmp := addmod(tmp, r, R_MOD)
      tmp := addmod(tmp, gamma, R_MOD)
      permutation := mulmod(permutation, tmp, R_MOD)
      tmp := mulmod(beta, calldataload(add(aproof, PROOF_S3_AT_ZETA_T)), R_MOD)
      tmp := addmod(tmp, o, R_MOD)
      tmp := addmod(tmp, gamma, R_MOD)
      permutation := mulmod(permutation, tmp, R_MOD)
      permutation := mulmod(permutation, calldataload(add(aproof, PROOF_Z_AT_ZETA_T_OMEGA)), R_MOD)

      tmp := mulmod(beta, zeta_t, R_MOD)
      tmp := addmod(tmp, l, R_MOD)
      tmp := addmod(tmp, gamma, R_MOD)
      let tmp2 := mulmod(beta, u_zeta_t, R_MOD)
      tmp2 := addmod(tmp2, r, R_MOD)
      tmp2 := addmod(tmp2, gamma, R_MOD)
      tmp := mulmod(tmp2, tmp, R_MOD)
      tmp2 := mulmod(beta, uu_zeta_t, R_MOD)
      tmp2 := addmod(tmp2, o, R_MOD)
      tmp2 := addmod(tmp2, gamma, R_MOD)
      tmp := mulmod(tmp, tmp2, R_MOD)
      tmp := mulmod(tmp, calldataload(add(aproof, PROOF_Z_AT_ZETA_T)), R_MOD)
      tmp := sub(R_MOD, tmp)
      permutation := addmod(permutation, tmp, R_MOD)
    }

    function check_z_start_at_one(aproof)->start_at_one {
      let state := mload(0x40)
      let z := calldataload(add(aproof, PROOF_Z_AT_ZETA_T))
      start_at_one := addmod(z, R_MOD_MINUS_ONE, R_MOD)
      let l0 := mload(add(state, STATE_LAGRANGE_0_AT_ZETA_T))
      start_at_one := mulmod(l0, start_at_one, R_MOD)
    }

	}

	return true;

  }
}
`

// MarshalSolidity convert  s a proof to a byte array that can be used in a
// Solidity contract.
func (proof *Proof) MarshalSolidity() []byte {

	res := make([]byte, 0, 1024)

	// uint256 lro_entangled_com_x;
	// uint256 lro_entangled_com_y;
	var tmp64 [64]byte
	tmp64 = proof.LROEntangled.RawBytes()
	res = append(res, tmp64[:]...)

	// uint256 Z non entangled
	tmp64 = proof.Z.RawBytes()
	res = append(res, tmp64[:]...)

	// uint256 Z entangled
	tmp64 = proof.ZEntangled.RawBytes()
	res = append(res, tmp64[:]...)

	// H entangled
	tmp64 = proof.HEntangled.RawBytes()
	res = append(res, tmp64[:]...)

	// BSB commitments
	for i := 0; i < len(proof.BsbComEntangled); i++ {
		tmp64 = proof.BsbComEntangled[i].RawBytes()
		res = append(res, tmp64[:]...)
	}

	// at this stage we serialise the fflonk proof

	// claimed values of (in that order):
	// ql, qr, qm, qo, qkIncomplete, s1, s2, s3, qcp_i, l, r, o, z, h1, h2, h3, bsb_i at ζ
	// z at ωζ
	var tmp32 [32]byte
	nbPolynomials := number_polynomials + 2*len(proof.BsbComEntangled)
	for i := 0; i < nbPolynomials; i++ { // -> there are some extra values that might be zero, we don't serialise them
		tmp32 = proof.BatchOpeningProof.ClaimedValues[0][i][0].Bytes()
		res = append(res, tmp32[:]...)
	}
	tmp32 = proof.BatchOpeningProof.ClaimedValues[1][0][0].Bytes()
	res = append(res, tmp32[:]...)

	// shplonk.W
	tmp64 = proof.BatchOpeningProof.SOpeningProof.W.RawBytes()
	res = append(res, tmp64[:]...)

	// shplonk.WPrime
	tmp64 = proof.BatchOpeningProof.SOpeningProof.WPrime.RawBytes()
	res = append(res, tmp64[:]...)

	// shplonk.ClaimedValues
	for i := 0; i < len(proof.BatchOpeningProof.SOpeningProof.ClaimedValues[0]); i++ {
		tmp32 = proof.BatchOpeningProof.SOpeningProof.ClaimedValues[0][i].Bytes()
		res = append(res, tmp32[:]...)
	}
	tmp32 = proof.BatchOpeningProof.SOpeningProof.ClaimedValues[1][0].Bytes()
	res = append(res, tmp32[:]...)

	return res
}
