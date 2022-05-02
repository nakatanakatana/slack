package slack

import (
	"context"
	"encoding/json"
)

type (
	WorkflowStepCompletedRequest struct {
		WorkflowStepExecuteID string             `json:"workflow_step_execute_id"`
		Outputs               *map[string]string `json:"outputs"`
	}

	WorkflowStepFailedRequest struct {
		WorkflowStepExecuteID string `json:"workflow_step_execute_id"`
		Error                 struct {
			Message string `json:"message"`
		} `json:"error"`
	}
)

func (api *Client) WorkflowStepCompleted(ctx context.Context, workflowStepExecuteID string, outputs *map[string]string) error {
	// More information: https://api.slack.com/methods/workflows.stepCompleted
	r := WorkflowStepCompletedRequest{
		WorkflowStepExecuteID: workflowStepExecuteID,
	}
	if outputs != nil {
		r.Outputs = outputs
	}

	endpoint := api.endpoint + "workflows.stepCompleted"
	jsonData, err := json.Marshal(r)
	if err != nil {
		return err
	}

	response := &SlackResponse{}
	if err := postJSON(ctx, api.httpclient, endpoint, api.token, jsonData, response, api); err != nil {
		return err
	}

	if !response.Ok {
		return response.Err()
	}

	return nil
}

func (api *Client) WorkflowStepFailed(ctx context.Context, workflowStepExecuteID string, errorMessage string) error {
	// More information: https://api.slack.com/methods/workflows.stepFailed
	r := WorkflowStepFailedRequest{
		WorkflowStepExecuteID: workflowStepExecuteID,
	}
	r.Error.Message = errorMessage

	endpoint := api.endpoint + "workflows.stepFailed"
	jsonData, err := json.Marshal(r)
	if err != nil {
		return err
	}

	response := &SlackResponse{}
	if err := postJSON(ctx, api.httpclient, endpoint, api.token, jsonData, response, api); err != nil {
		return err
	}

	if !response.Ok {
		return response.Err()
	}

	return nil
}
