// Code generated by protoc-gen-validator. DO NOT EDIT.
// source: mail.proto

package pb

import (
	fmt "fmt"
	validator "gitlab.gg.com/game_framework/commons-go/validator"
)

func (this *MailSystem) Validate() error {
	return nil
}
func (this *MailSystem_EnvValue) Validate() error {
	return nil
}
func (this *MailSystem_EnvWrap) Validate() error {
	if nil == this.Val {
		return validator.FieldError("Val", fmt.Errorf("message must exist"))
	}
	if this.Val != nil {
		if err := validator.CallValidatorIfExists(this.Val); err != nil {
			return validator.FieldError("Val", err)
		}
	}
	if !(this.Duration > -1) {
		return validator.FieldError("Duration", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Duration))
	}
	return nil
}
func (this *MailSystem_PlayerInfo) Validate() error {
	if !(this.Id > 0) {
		return validator.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	if this.Name == "" {
		return validator.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *MailSystem_SystemInfo) Validate() error {
	if this.Name == "" {
		return validator.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *MailSystem_Box) Validate() error {
	for _, item := range this.Mails {
		if item != nil {
			if err := validator.CallValidatorIfExists(item); err != nil {
				return validator.FieldError("Mails", err)
			}
		}
	}
	for _, item := range this.GMails {
		if item != nil {
			if err := validator.CallValidatorIfExists(item); err != nil {
				return validator.FieldError("GMails", err)
			}
		}
	}
	return nil
}
func (this *MailSystem_Mail) Validate() error {
	if !(this.ReceiverId > 0) {
		return validator.FieldError("ReceiverId", fmt.Errorf(`value '%v' must be greater than '0'`, this.ReceiverId))
	}
	if this.Sender != nil {
		if err := validator.CallValidatorIfExists(this.Sender); err != nil {
			return validator.FieldError("Sender", err)
		}
	}
	if this.Title == "" {
		return validator.FieldError("Title", fmt.Errorf(`value '%v' must not be an empty string`, this.Title))
	}
	for _, item := range this.AttachArray {
		if item != nil {
			if err := validator.CallValidatorIfExists(item); err != nil {
				return validator.FieldError("AttachArray", err)
			}
		}
	}
	return nil
}
func (this *MailSystem_GlobalMail) Validate() error {
	if this.Sender != nil {
		if err := validator.CallValidatorIfExists(this.Sender); err != nil {
			return validator.FieldError("Sender", err)
		}
	}
	if this.Title == "" {
		return validator.FieldError("Title", fmt.Errorf(`value '%v' must not be an empty string`, this.Title))
	}
	for _, item := range this.AttachArray {
		if item != nil {
			if err := validator.CallValidatorIfExists(item); err != nil {
				return validator.FieldError("AttachArray", err)
			}
		}
	}
	return nil
}
func (this *MailSystem_Attachment) Validate() error {
	return nil
}
func (this *MailSystem_SendRequest) Validate() error {
	if nil == this.Mail {
		return validator.FieldError("Mail", fmt.Errorf("message must exist"))
	}
	if this.Mail != nil {
		if err := validator.CallValidatorIfExists(this.Mail); err != nil {
			return validator.FieldError("Mail", err)
		}
	}
	return nil
}
func (this *MailSystem_SendResponse) Validate() error {
	if this.Mail != nil {
		if err := validator.CallValidatorIfExists(this.Mail); err != nil {
			return validator.FieldError("Mail", err)
		}
	}
	return nil
}
func (this *MailSystem_PullRequest) Validate() error {
	if !(this.PlayerId > 0) {
		return validator.FieldError("PlayerId", fmt.Errorf(`value '%v' must be greater than '0'`, this.PlayerId))
	}
	if !(this.MailOffset > -1) {
		return validator.FieldError("MailOffset", fmt.Errorf(`value '%v' must be greater than '-1'`, this.MailOffset))
	}
	for _, item := range this.Env {
		if item != nil {
			if err := validator.CallValidatorIfExists(item); err != nil {
				return validator.FieldError("Env", err)
			}
		}
	}
	return nil
}
func (this *MailSystem_PullResponse) Validate() error {
	for _, item := range this.Mails {
		if item != nil {
			if err := validator.CallValidatorIfExists(item); err != nil {
				return validator.FieldError("Mails", err)
			}
		}
	}
	return nil
}
func (this *MailSystem_ReadRequest) Validate() error {
	if !(this.PlayerId > 0) {
		return validator.FieldError("PlayerId", fmt.Errorf(`value '%v' must be greater than '0'`, this.PlayerId))
	}
	if len(this.Ids) < 1 {
		return validator.FieldError("Ids", fmt.Errorf(`value '%v' must contain at least 1 elements`, this.Ids))
	}
	if !(this.Op >= 0) {
		return validator.FieldError("Op", fmt.Errorf(`value '%v' must be greater than '0'`, this.Op))
	}
	if !(this.Op <= 5) {
		return validator.FieldError("Op", fmt.Errorf(`value '%v' must be less than '0'`, this.Op))
	}
	return nil
}
func (this *MailSystem_ReadResponse) Validate() error {
	for _, item := range this.Mails {
		if item != nil {
			if err := validator.CallValidatorIfExists(item); err != nil {
				return validator.FieldError("Mails", err)
			}
		}
	}
	return nil
}
func (this *MailSystem_AttachRequest) Validate() error {
	if !(this.PlayerId > 0) {
		return validator.FieldError("PlayerId", fmt.Errorf(`value '%v' must be greater than '0'`, this.PlayerId))
	}
	if !(this.MailId > -1) {
		return validator.FieldError("MailId", fmt.Errorf(`value '%v' must be greater than '-1'`, this.MailId))
	}
	return nil
}
func (this *MailSystem_AttachResponse) Validate() error {
	for _, item := range this.Packs {
		if item != nil {
			if err := validator.CallValidatorIfExists(item); err != nil {
				return validator.FieldError("Packs", err)
			}
		}
	}
	return nil
}
func (this *MailSystem_AttachmentPack) Validate() error {
	for _, item := range this.AttachArray {
		if item != nil {
			if err := validator.CallValidatorIfExists(item); err != nil {
				return validator.FieldError("AttachArray", err)
			}
		}
	}
	return nil
}
func (this *MailSystem_SubmitRequest) Validate() error {
	if nil == this.Gmail {
		return validator.FieldError("Gmail", fmt.Errorf("message must exist"))
	}
	if this.Gmail != nil {
		if err := validator.CallValidatorIfExists(this.Gmail); err != nil {
			return validator.FieldError("Gmail", err)
		}
	}
	for _, item := range this.TestParam {
		if item != nil {
			if err := validator.CallValidatorIfExists(item); err != nil {
				return validator.FieldError("TestParam", err)
			}
		}
	}
	return nil
}
func (this *MailSystem_SubmitResponse) Validate() error {
	return nil
}
func (this *MailSystem_DeleteGlobalMailRequest) Validate() error {
	if len(this.Ids) < 1 {
		return validator.FieldError("Ids", fmt.Errorf(`value '%v' must contain at least 1 elements`, this.Ids))
	}
	return nil
}
func (this *MailSystem_DeleteGlobalMailResponse) Validate() error {
	return nil
}
func (this *MailSystem_RemoveRequest) Validate() error {
	if !(this.PlayerId > 0) {
		return validator.FieldError("PlayerId", fmt.Errorf(`value '%v' must be greater than '0'`, this.PlayerId))
	}
	if !(this.Policy >= 0) {
		return validator.FieldError("Policy", fmt.Errorf(`value '%v' must be greater than '0'`, this.Policy))
	}
	if !(this.Policy <= 3) {
		return validator.FieldError("Policy", fmt.Errorf(`value '%v' must be less than '0'`, this.Policy))
	}
	return nil
}
func (this *MailSystem_RemoveResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *MailSystem_NewGlobalEnvRequest) Validate() error {
	for _, item := range this.Env {
		if item != nil {
			if err := validator.CallValidatorIfExists(item); err != nil {
				return validator.FieldError("Env", err)
			}
		}
	}
	return nil
}
func (this *MailSystem_NewGlobalEnvResponse) Validate() error {
	for _, item := range this.Env {
		if item != nil {
			if err := validator.CallValidatorIfExists(item); err != nil {
				return validator.FieldError("Env", err)
			}
		}
	}
	return nil
}
func (this *MailSystem_RemoveGlobalEnvRequest) Validate() error {
	if len(this.Keys) < 1 {
		return validator.FieldError("Keys", fmt.Errorf(`value '%v' must contain at least 1 elements`, this.Keys))
	}
	return nil
}
func (this *MailSystem_RemoveGlobalEnvResponse) Validate() error {
	if !(this.Status >= 0) {
		return validator.FieldError("Status", fmt.Errorf(`value '%v' must be greater than '0'`, this.Status))
	}
	if !(this.Status <= 1) {
		return validator.FieldError("Status", fmt.Errorf(`value '%v' must be less than '0'`, this.Status))
	}
	for _, item := range this.Env {
		if item != nil {
			if err := validator.CallValidatorIfExists(item); err != nil {
				return validator.FieldError("Env", err)
			}
		}
	}
	return nil
}
