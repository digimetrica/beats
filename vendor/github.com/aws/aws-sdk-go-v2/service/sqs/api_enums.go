// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

type MessageSystemAttributeName string

// Enum values for MessageSystemAttributeName
const (
	MessageSystemAttributeNameSenderId                         MessageSystemAttributeName = "SenderId"
	MessageSystemAttributeNameSentTimestamp                    MessageSystemAttributeName = "SentTimestamp"
	MessageSystemAttributeNameApproximateReceiveCount          MessageSystemAttributeName = "ApproximateReceiveCount"
	MessageSystemAttributeNameApproximateFirstReceiveTimestamp MessageSystemAttributeName = "ApproximateFirstReceiveTimestamp"
	MessageSystemAttributeNameSequenceNumber                   MessageSystemAttributeName = "SequenceNumber"
	MessageSystemAttributeNameMessageDeduplicationId           MessageSystemAttributeName = "MessageDeduplicationId"
	MessageSystemAttributeNameMessageGroupId                   MessageSystemAttributeName = "MessageGroupId"
)

func (enum MessageSystemAttributeName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MessageSystemAttributeName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type QueueAttributeName string

// Enum values for QueueAttributeName
const (
	QueueAttributeNameAll                                   QueueAttributeName = "All"
	QueueAttributeNamePolicy                                QueueAttributeName = "Policy"
	QueueAttributeNameVisibilityTimeout                     QueueAttributeName = "VisibilityTimeout"
	QueueAttributeNameMaximumMessageSize                    QueueAttributeName = "MaximumMessageSize"
	QueueAttributeNameMessageRetentionPeriod                QueueAttributeName = "MessageRetentionPeriod"
	QueueAttributeNameApproximateNumberOfMessages           QueueAttributeName = "ApproximateNumberOfMessages"
	QueueAttributeNameApproximateNumberOfMessagesNotVisible QueueAttributeName = "ApproximateNumberOfMessagesNotVisible"
	QueueAttributeNameCreatedTimestamp                      QueueAttributeName = "CreatedTimestamp"
	QueueAttributeNameLastModifiedTimestamp                 QueueAttributeName = "LastModifiedTimestamp"
	QueueAttributeNameQueueArn                              QueueAttributeName = "QueueArn"
	QueueAttributeNameApproximateNumberOfMessagesDelayed    QueueAttributeName = "ApproximateNumberOfMessagesDelayed"
	QueueAttributeNameDelaySeconds                          QueueAttributeName = "DelaySeconds"
	QueueAttributeNameReceiveMessageWaitTimeSeconds         QueueAttributeName = "ReceiveMessageWaitTimeSeconds"
	QueueAttributeNameRedrivePolicy                         QueueAttributeName = "RedrivePolicy"
	QueueAttributeNameFifoQueue                             QueueAttributeName = "FifoQueue"
	QueueAttributeNameContentBasedDeduplication             QueueAttributeName = "ContentBasedDeduplication"
	QueueAttributeNameKmsMasterKeyId                        QueueAttributeName = "KmsMasterKeyId"
	QueueAttributeNameKmsDataKeyReusePeriodSeconds          QueueAttributeName = "KmsDataKeyReusePeriodSeconds"
)

func (enum QueueAttributeName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum QueueAttributeName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
