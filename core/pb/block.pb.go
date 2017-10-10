// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

/*
	Package corepb is a generated protocol buffer package.

	It is generated from these files:
		block.proto

	It has these top-level messages:
		Transaction
		BlockHeader
		Block
*/
package corepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Transaction struct {
	Hash      []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	From      []byte `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To        []byte `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Value     uint64 `protobuf:"varint,4,opt,name=value,proto3" json:"value,omitempty"`
	Nonce     uint64 `protobuf:"varint,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Timestamp int64  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data      []byte `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	Alg       uint32 `protobuf:"varint,8,opt,name=alg,proto3" json:"alg,omitempty"`
	Sign      []byte `protobuf:"bytes,9,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{0} }

func (m *Transaction) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Transaction) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Transaction) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Transaction) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Transaction) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Transaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Transaction) GetAlg() uint32 {
	if m != nil {
		return m.Alg
	}
	return 0
}

func (m *Transaction) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type BlockHeader struct {
	Hash       []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ParentHash []byte `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	StateRoot  []byte `protobuf:"bytes,3,opt,name=stateRoot,proto3" json:"stateRoot,omitempty"`
	Nonce      uint64 `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Coinbase   []byte `protobuf:"bytes,5,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	Timestamp  int64  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{1} }

func (m *BlockHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockHeader) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *BlockHeader) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *BlockHeader) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *BlockHeader) GetCoinbase() []byte {
	if m != nil {
		return m.Coinbase
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Block struct {
	Header       *BlockHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{2} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterType((*Transaction)(nil), "corepb.Transaction")
	proto.RegisterType((*BlockHeader)(nil), "corepb.BlockHeader")
	proto.RegisterType((*Block)(nil), "corepb.Block")
}
func (m *Transaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transaction) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	if len(m.From) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBlock(dAtA, i, uint64(len(m.From)))
		i += copy(dAtA[i:], m.From)
	}
	if len(m.To) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBlock(dAtA, i, uint64(len(m.To)))
		i += copy(dAtA[i:], m.To)
	}
	if m.Value != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintBlock(dAtA, i, uint64(m.Value))
	}
	if m.Nonce != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintBlock(dAtA, i, uint64(m.Nonce))
	}
	if m.Timestamp != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintBlock(dAtA, i, uint64(m.Timestamp))
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.Alg != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintBlock(dAtA, i, uint64(m.Alg))
	}
	if len(m.Sign) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Sign)))
		i += copy(dAtA[i:], m.Sign)
	}
	return i, nil
}

func (m *BlockHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	if len(m.ParentHash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBlock(dAtA, i, uint64(len(m.ParentHash)))
		i += copy(dAtA[i:], m.ParentHash)
	}
	if len(m.StateRoot) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBlock(dAtA, i, uint64(len(m.StateRoot)))
		i += copy(dAtA[i:], m.StateRoot)
	}
	if m.Nonce != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintBlock(dAtA, i, uint64(m.Nonce))
	}
	if len(m.Coinbase) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Coinbase)))
		i += copy(dAtA[i:], m.Coinbase)
	}
	if m.Timestamp != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintBlock(dAtA, i, uint64(m.Timestamp))
	}
	return i, nil
}

func (m *Block) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Block) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBlock(dAtA, i, uint64(m.Header.Size()))
		n1, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Transactions) > 0 {
		for _, msg := range m.Transactions {
			dAtA[i] = 0x12
			i++
			i = encodeVarintBlock(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Block(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Block(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Transaction) Size() (n int) {
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovBlock(uint64(m.Value))
	}
	if m.Nonce != 0 {
		n += 1 + sovBlock(uint64(m.Nonce))
	}
	if m.Timestamp != 0 {
		n += 1 + sovBlock(uint64(m.Timestamp))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Alg != 0 {
		n += 1 + sovBlock(uint64(m.Alg))
	}
	l = len(m.Sign)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	return n
}

func (m *BlockHeader) Size() (n int) {
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.ParentHash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.StateRoot)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovBlock(uint64(m.Nonce))
	}
	l = len(m.Coinbase)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovBlock(uint64(m.Timestamp))
	}
	return n
}

func (m *Block) Size() (n int) {
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if len(m.Transactions) > 0 {
		for _, e := range m.Transactions {
			l = e.Size()
			n += 1 + l + sovBlock(uint64(l))
		}
	}
	return n
}

func sovBlock(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Transaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Transaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = append(m.From[:0], dAtA[iNdEx:postIndex]...)
			if m.From == nil {
				m.From = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = append(m.To[:0], dAtA[iNdEx:postIndex]...)
			if m.To == nil {
				m.To = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alg", wireType)
			}
			m.Alg = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Alg |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sign", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sign = append(m.Sign[:0], dAtA[iNdEx:postIndex]...)
			if m.Sign == nil {
				m.Sign = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentHash = append(m.ParentHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ParentHash == nil {
				m.ParentHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateRoot = append(m.StateRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.StateRoot == nil {
				m.StateRoot = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coinbase", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coinbase = append(m.Coinbase[:0], dAtA[iNdEx:postIndex]...)
			if m.Coinbase == nil {
				m.Coinbase = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Block) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Block: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &BlockHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transactions = append(m.Transactions, &Transaction{})
			if err := m.Transactions[len(m.Transactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthBlock
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBlock
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBlock(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBlock = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("block.proto", fileDescriptorBlock) }

var fileDescriptorBlock = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x4e, 0xeb, 0x30,
	0x10, 0x86, 0x9f, 0x93, 0x34, 0xaf, 0x9d, 0x14, 0x54, 0x19, 0x16, 0x16, 0x42, 0x51, 0xd4, 0x55,
	0x24, 0xa4, 0x2e, 0x60, 0xc1, 0xbe, 0xab, 0xae, 0x2d, 0x2e, 0xe0, 0xa4, 0xa6, 0x8d, 0x68, 0xec,
	0xc8, 0x1e, 0x38, 0x0b, 0xb7, 0xe0, 0x1a, 0xb0, 0xe3, 0x08, 0x28, 0x5c, 0x04, 0xd9, 0xae, 0x48,
	0x90, 0x10, 0xbb, 0x7f, 0xbe, 0x4c, 0x34, 0xf3, 0x69, 0x0c, 0x59, 0x75, 0xd0, 0xf5, 0xc3, 0xaa,
	0x33, 0x1a, 0x35, 0x4d, 0x6b, 0x6d, 0x64, 0x57, 0x2d, 0xdf, 0x08, 0x64, 0x77, 0x46, 0x28, 0x2b,
	0x6a, 0x6c, 0xb4, 0xa2, 0x14, 0x92, 0xbd, 0xb0, 0x7b, 0x46, 0x0a, 0x52, 0xce, 0xb9, 0xcf, 0x8e,
	0xdd, 0x1b, 0xdd, 0xb2, 0x28, 0x30, 0x97, 0xe9, 0x29, 0x44, 0xa8, 0x59, 0xec, 0x49, 0x84, 0x9a,
	0x9e, 0xc3, 0xe4, 0x49, 0x1c, 0x1e, 0x25, 0x4b, 0x0a, 0x52, 0x26, 0x3c, 0x14, 0x8e, 0x2a, 0xad,
	0x6a, 0xc9, 0x26, 0x81, 0xfa, 0x82, 0x5e, 0xc2, 0x0c, 0x9b, 0x56, 0x5a, 0x14, 0x6d, 0xc7, 0xd2,
	0x82, 0x94, 0x31, 0x1f, 0x80, 0x9b, 0xb6, 0x15, 0x28, 0xd8, 0xff, 0x30, 0xcd, 0x65, 0xba, 0x80,
	0x58, 0x1c, 0x76, 0x6c, 0x5a, 0x90, 0xf2, 0x84, 0xbb, 0xe8, 0xba, 0x6c, 0xb3, 0x53, 0x6c, 0x16,
	0xba, 0x5c, 0x5e, 0xbe, 0x10, 0xc8, 0xd6, 0xce, 0x71, 0x23, 0xc5, 0x56, 0x9a, 0x5f, 0x5d, 0x72,
	0x80, 0x4e, 0x18, 0xa9, 0x70, 0xe3, 0xbe, 0x04, 0xa3, 0x11, 0x71, 0xbb, 0x59, 0x14, 0x28, 0xb9,
	0xd6, 0x78, 0xd4, 0x1b, 0xc0, 0xe0, 0x93, 0x8c, 0x7d, 0x2e, 0x60, 0x5a, 0xeb, 0x46, 0x55, 0xc2,
	0x06, 0xd1, 0x39, 0xff, 0xae, 0xff, 0x76, 0x5d, 0xb6, 0x30, 0xf1, 0x0b, 0xd3, 0x2b, 0x48, 0xf7,
	0x7e, 0x69, 0xbf, 0x6c, 0x76, 0x7d, 0xb6, 0x0a, 0xf7, 0x59, 0x8d, 0x7c, 0xf8, 0xb1, 0x85, 0xde,
	0xc2, 0x1c, 0x87, 0x93, 0x59, 0x16, 0x15, 0xf1, 0xf8, 0x97, 0xd1, 0x39, 0xf9, 0x8f, 0xc6, 0xf5,
	0xe2, 0xb5, 0xcf, 0xc9, 0x7b, 0x9f, 0x93, 0x8f, 0x3e, 0x27, 0xcf, 0x9f, 0xf9, 0xbf, 0x2a, 0xf5,
	0xaf, 0xe1, 0xe6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x69, 0x86, 0xea, 0x1c, 0x02, 0x00, 0x00,
}
