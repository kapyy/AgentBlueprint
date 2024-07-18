package py_comm_client

import (
	"errors"
	message "golang-client/message/protoData"
	"golang-client/modules/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	PyClient message.APMServiceClient
	PyConn   *grpc.ClientConn
)

func ConnectRPCClient(addr string) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	var err error
	PyConn, err = grpc.NewClient(addr, opts...)
	if err != nil {
		logger.GetLogger().WithField("func", "ConnectRPCClient").Errorf("ConnectRPCClient error: %v,with Addr %s", err, addr)
		return err
	}
	PyClient = message.NewAPMServiceClient(PyConn)

	return nil
}

func SendMainServiceRequest(requestMessage *message.NodeData) (*message.ServiceResponse, error) {
	log := logger.GetLogger().WithField("func", "SendMainServiceRequest")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	srvsrqst := message.MainServicerRequest{
		MessageId: 1,
		Data:      requestMessage,
	}
	if PyClient == nil {
		log.Errorf("PyClient is Not Connected")
		return nil, errors.New("PyClient is Not Connected")
	}
	return PyClient.MainServiceRequest(ctx, &srvsrqst)

}
func SendSubordinateServiceRequest(function_id uint64, data_id uint64, data []byte) (*message.ServiceResponse, error) {
	log := logger.GetLogger().WithField("func", "SendSubordinateServiceRequest")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	srvsrqst := message.SubordinateServicerRequest{
		MessageId: function_id,
		DataType:  data_id,
		RqstData:  data,
	}
	if PyClient == nil {
		log.Errorf("PyClient is Not Connected")
		return nil, errors.New("PyClient is Not Connected")
	}
	response, err := PyClient.SubordinateServiceRequest(ctx, &srvsrqst)
	if err != nil {
		log.Errorf("SendSubordinateServiceRequest error: %v", err)
	}
	return response, nil
}
