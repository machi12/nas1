// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/machi12/nas/nasType"
)

type PDUSessionReleaseCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONRELEASECOMMANDMessageIdentity
	nasType.Cause5GSM
	*nasType.BackoffTimerValue
	*nasType.EAPMessage
	*nasType.CongestionReattemptIndicator5GSM
	*nasType.ExtendedProtocolConfigurationOptions
}

func NewPDUSessionReleaseCommand(iei uint8) (pDUSessionReleaseCommand *PDUSessionReleaseCommand) {
	pDUSessionReleaseCommand = &PDUSessionReleaseCommand{}
	return pDUSessionReleaseCommand
}

const (
	PDUSessionReleaseCommandBackoffTimerValueType                    uint8 = 0x37
	PDUSessionReleaseCommandEAPMessageType                           uint8 = 0x78
	PDUSessionReleaseCommandCongestionReattemptIndicator5GSMType     uint8 = 0x61
	PDUSessionReleaseCommandExtendedProtocolConfigurationOptionsType uint8 = 0x7B
)

func (a *PDUSessionReleaseCommand) EncodePDUSessionReleaseCommand(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionID.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/PDUSessionID): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PTI.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/PTI): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PDUSESSIONRELEASECOMMANDMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/PDUSESSIONRELEASECOMMANDMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.Cause5GSM.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/Cause5GSM): %w", err)
	}
	if a.BackoffTimerValue != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/BackoffTimerValue): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/BackoffTimerValue): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.Octet); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/BackoffTimerValue): %w", err)
		}
	}
	if a.EAPMessage != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/EAPMessage): %w", err)
		}
	}
	if a.CongestionReattemptIndicator5GSM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.CongestionReattemptIndicator5GSM.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/CongestionReattemptIndicator5GSM): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.CongestionReattemptIndicator5GSM.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/CongestionReattemptIndicator5GSM): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.CongestionReattemptIndicator5GSM.Octet); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/CongestionReattemptIndicator5GSM): %w", err)
		}
	}
	if a.ExtendedProtocolConfigurationOptions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/ExtendedProtocolConfigurationOptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/ExtendedProtocolConfigurationOptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseCommand/ExtendedProtocolConfigurationOptions): %w", err)
		}
	}
	return nil
}

func (a *PDUSessionReleaseCommand) DecodePDUSessionReleaseCommand(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionReleaseCommand/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionReleaseCommand/PDUSessionID): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionReleaseCommand/PTI): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONRELEASECOMMANDMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionReleaseCommand/PDUSESSIONRELEASECOMMANDMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionReleaseCommand/Cause5GSM): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (PDUSessionReleaseCommand/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case PDUSessionReleaseCommandBackoffTimerValueType:
			a.BackoffTimerValue = nasType.NewBackoffTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionReleaseCommand/BackoffTimerValue): %w", err)
			}
			if a.BackoffTimerValue.Len != 1 {
				return fmt.Errorf("invalid ie length (PDUSessionReleaseCommand/BackoffTimerValue): %d", a.BackoffTimerValue.Len)
			}
			a.BackoffTimerValue.SetLen(a.BackoffTimerValue.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionReleaseCommand/BackoffTimerValue): %w", err)
			}
		case PDUSessionReleaseCommandEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionReleaseCommand/EAPMessage): %w", err)
			}
			if a.EAPMessage.Len < 4 || a.EAPMessage.Len > 1500 {
				return fmt.Errorf("invalid ie length (PDUSessionReleaseCommand/EAPMessage): %d", a.EAPMessage.Len)
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionReleaseCommand/EAPMessage): %w", err)
			}
		case PDUSessionReleaseCommandCongestionReattemptIndicator5GSMType:
			a.CongestionReattemptIndicator5GSM = nasType.NewCongestionReattemptIndicator5GSM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.CongestionReattemptIndicator5GSM.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionReleaseCommand/CongestionReattemptIndicator5GSM): %w", err)
			}
			if a.CongestionReattemptIndicator5GSM.Len != 1 {
				return fmt.Errorf("invalid ie length (PDUSessionReleaseCommand/CongestionReattemptIndicator5GSM): %d", a.CongestionReattemptIndicator5GSM.Len)
			}
			a.CongestionReattemptIndicator5GSM.SetLen(a.CongestionReattemptIndicator5GSM.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.CongestionReattemptIndicator5GSM.Octet); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionReleaseCommand/CongestionReattemptIndicator5GSM): %w", err)
			}
		case PDUSessionReleaseCommandExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionReleaseCommand/ExtendedProtocolConfigurationOptions): %w", err)
			}
			if a.ExtendedProtocolConfigurationOptions.Len < 1 {
				return fmt.Errorf("invalid ie length (PDUSessionReleaseCommand/ExtendedProtocolConfigurationOptions): %d", a.ExtendedProtocolConfigurationOptions.Len)
			}
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionReleaseCommand/ExtendedProtocolConfigurationOptions): %w", err)
			}
		default:
		}
	}
	return nil
}
