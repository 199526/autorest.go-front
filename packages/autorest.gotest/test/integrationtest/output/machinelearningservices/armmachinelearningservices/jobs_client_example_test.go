//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearningservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/AutoMLJob/list.json
func ExampleJobsClient_NewListPager_list() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("test-rg", "my-aml-workspace", &armmachinelearningservices.JobsClientListOptions{Skip: nil,
		JobType:      nil,
		Tag:          nil,
		ListViewType: nil,
		Scheduled:    nil,
		ScheduleID:   nil,
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/CommandJob/list.json
func ExampleJobsClient_NewListPager_list() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("test-rg", "my-aml-workspace", &armmachinelearningservices.JobsClientListOptions{Skip: nil,
		JobType:      to.Ptr("string"),
		Tag:          to.Ptr("string"),
		ListViewType: nil,
		Scheduled:    nil,
		ScheduleID:   nil,
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/PipelineJob/list.json
func ExampleJobsClient_NewListPager_list() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("test-rg", "my-aml-workspace", &armmachinelearningservices.JobsClientListOptions{Skip: nil,
		JobType:      to.Ptr("string"),
		Tag:          to.Ptr("string"),
		ListViewType: nil,
		Scheduled:    nil,
		ScheduleID:   nil,
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/SweepJob/list.json
func ExampleJobsClient_NewListPager_list() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("test-rg", "my-aml-workspace", &armmachinelearningservices.JobsClientListOptions{Skip: nil,
		JobType:      to.Ptr("string"),
		Tag:          to.Ptr("string"),
		ListViewType: nil,
		Scheduled:    nil,
		ScheduleID:   nil,
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/delete.json
func ExampleJobsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "test-rg", "my-aml-workspace", "string", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/AutoMLJob/get.json
func ExampleJobsClient_Get_get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "test-rg", "my-aml-workspace", "string", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/CommandJob/get.json
func ExampleJobsClient_Get_get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "test-rg", "my-aml-workspace", "string", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/PipelineJob/get.json
func ExampleJobsClient_Get_get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "test-rg", "my-aml-workspace", "string", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/SweepJob/get.json
func ExampleJobsClient_Get_get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "test-rg", "my-aml-workspace", "string", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/AutoMLJob/createOrUpdate.json
func ExampleJobsClient_CreateOrUpdate_createOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "string", armmachinelearningservices.JobBaseData{
		Properties: &armmachinelearningservices.AutoMLJob{
			Description: to.Ptr("string"),
			Properties: map[string]*string{
				"string": to.Ptr("string"),
			},
			Tags: map[string]*string{
				"string": to.Ptr("string"),
			},
			ComputeID:      to.Ptr("string"),
			DisplayName:    to.Ptr("string"),
			ExperimentName: to.Ptr("string"),
			Identity: &armmachinelearningservices.AmlToken{
				IdentityType: to.Ptr(armmachinelearningservices.IdentityConfigurationTypeAMLToken),
			},
			IsArchived: to.Ptr(false),
			JobType:    to.Ptr(armmachinelearningservices.JobTypeAutoML),
			Schedule: &armmachinelearningservices.CronSchedule{
				EndTime:        to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t }()),
				ScheduleStatus: to.Ptr(armmachinelearningservices.ScheduleStatusDisabled),
				ScheduleType:   to.Ptr(armmachinelearningservices.ScheduleTypeCron),
				StartTime:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t }()),
				TimeZone:       to.Ptr("string"),
				Expression:     to.Ptr("string"),
			},
			Services: map[string]*armmachinelearningservices.JobService{
				"string": &armmachinelearningservices.JobService{
					Endpoint:       to.Ptr("string"),
					JobServiceType: to.Ptr("string"),
					Port:           to.Ptr[int32](1),
					Properties: map[string]*string{
						"string": to.Ptr("string"),
					},
				},
			},
			EnvironmentID: to.Ptr("string"),
			EnvironmentVariables: map[string]*string{
				"string": to.Ptr("string"),
			},
			Outputs: map[string]armmachinelearningservices.JobOutputClassification{
				"string": &armmachinelearningservices.URIFileJobOutput{
					Mode:          to.Ptr(armmachinelearningservices.OutputDeliveryModeReadWriteMount),
					URI:           to.Ptr("string"),
					Description:   to.Ptr("string"),
					JobOutputType: to.Ptr(armmachinelearningservices.JobOutputTypeURIFile),
				},
			},
			Resources: &armmachinelearningservices.ResourceConfiguration{
				InstanceCount: to.Ptr[int32](1),
				InstanceType:  to.Ptr("string"),
				Properties: map[string]interface{}{
					"string": map[string]interface{}{
						"9bec0ab0-c62f-4fa9-a97c-7b24bbcc90ad": nil,
					},
				},
			},
			TaskDetails: &armmachinelearningservices.ImageClassification{
				TaskType: to.Ptr(armmachinelearningservices.TaskTypeImageClassification),
				DataSettings: &armmachinelearningservices.ImageVerticalDataSettings{
					TargetColumnName: to.Ptr("string"),
					TrainingData: &armmachinelearningservices.TrainingDataSettings{
						Data: &armmachinelearningservices.MLTableJobInput{
							URI:          to.Ptr("string"),
							JobInputType: to.Ptr(armmachinelearningservices.JobInputTypeMLTable),
						},
					},
				},
				LimitSettings: &armmachinelearningservices.ImageLimitSettings{
					MaxTrials: to.Ptr[int32](2),
				},
				ModelSettings: &armmachinelearningservices.ImageModelSettingsClassification{
					ValidationCropSize: to.Ptr[int32](2),
				},
				SearchSpace: []*armmachinelearningservices.ImageModelDistributionSettingsClassification{
					{
						ValidationCropSize: to.Ptr("choice(2, 360)"),
					}},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/CommandJob/createOrUpdate.json
func ExampleJobsClient_CreateOrUpdate_createOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "string", armmachinelearningservices.JobBaseData{
		Properties: &armmachinelearningservices.CommandJob{
			Description: to.Ptr("string"),
			Properties: map[string]*string{
				"string": to.Ptr("string"),
			},
			Tags: map[string]*string{
				"string": to.Ptr("string"),
			},
			ComputeID:      to.Ptr("string"),
			DisplayName:    to.Ptr("string"),
			ExperimentName: to.Ptr("string"),
			Identity: &armmachinelearningservices.AmlToken{
				IdentityType: to.Ptr(armmachinelearningservices.IdentityConfigurationTypeAMLToken),
			},
			JobType: to.Ptr(armmachinelearningservices.JobTypeCommand),
			Services: map[string]*armmachinelearningservices.JobService{
				"string": &armmachinelearningservices.JobService{
					Endpoint:       to.Ptr("string"),
					JobServiceType: to.Ptr("string"),
					Port:           to.Ptr[int32](1),
					Properties: map[string]*string{
						"string": to.Ptr("string"),
					},
				},
			},
			CodeID:  to.Ptr("string"),
			Command: to.Ptr("string"),
			Distribution: &armmachinelearningservices.TensorFlow{
				DistributionType:     to.Ptr(armmachinelearningservices.DistributionTypeTensorFlow),
				ParameterServerCount: to.Ptr[int32](1),
				WorkerCount:          to.Ptr[int32](1),
			},
			EnvironmentID: to.Ptr("string"),
			EnvironmentVariables: map[string]*string{
				"string": to.Ptr("string"),
			},
			Inputs: map[string]armmachinelearningservices.JobInputClassification{
				"string": &armmachinelearningservices.LiteralJobInput{
					Description:  to.Ptr("string"),
					JobInputType: to.Ptr(armmachinelearningservices.JobInputTypeLiteral),
					Value:        to.Ptr("string"),
				},
			},
			Limits: &armmachinelearningservices.CommandJobLimits{
				JobLimitsType: to.Ptr(armmachinelearningservices.JobLimitsTypeCommand),
				Timeout:       to.Ptr("PT5M"),
			},
			Outputs: map[string]armmachinelearningservices.JobOutputClassification{
				"string": &armmachinelearningservices.URIFileJobOutput{
					Mode:          to.Ptr(armmachinelearningservices.OutputDeliveryModeReadWriteMount),
					URI:           to.Ptr("string"),
					Description:   to.Ptr("string"),
					JobOutputType: to.Ptr(armmachinelearningservices.JobOutputTypeURIFile),
				},
			},
			Resources: &armmachinelearningservices.ResourceConfiguration{
				InstanceCount: to.Ptr[int32](1),
				InstanceType:  to.Ptr("string"),
				Properties: map[string]interface{}{
					"string": map[string]interface{}{
						"e6b6493e-7d5e-4db3-be1e-306ec641327e": nil,
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/PipelineJob/createOrUpdate.json
func ExampleJobsClient_CreateOrUpdate_createOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "string", armmachinelearningservices.JobBaseData{
		Properties: &armmachinelearningservices.PipelineJob{
			Description: to.Ptr("string"),
			Properties: map[string]*string{
				"string": to.Ptr("string"),
			},
			Tags: map[string]*string{
				"string": to.Ptr("string"),
			},
			ComputeID:      to.Ptr("string"),
			DisplayName:    to.Ptr("string"),
			ExperimentName: to.Ptr("string"),
			JobType:        to.Ptr(armmachinelearningservices.JobTypePipeline),
			Services: map[string]*armmachinelearningservices.JobService{
				"string": &armmachinelearningservices.JobService{
					Endpoint:       to.Ptr("string"),
					JobServiceType: to.Ptr("string"),
					Port:           to.Ptr[int32](1),
					Properties: map[string]*string{
						"string": to.Ptr("string"),
					},
				},
			},
			Inputs: map[string]armmachinelearningservices.JobInputClassification{
				"string": &armmachinelearningservices.LiteralJobInput{
					Description:  to.Ptr("string"),
					JobInputType: to.Ptr(armmachinelearningservices.JobInputTypeLiteral),
					Value:        to.Ptr("string"),
				},
			},
			Outputs: map[string]armmachinelearningservices.JobOutputClassification{
				"string": &armmachinelearningservices.URIFileJobOutput{
					Mode:          to.Ptr(armmachinelearningservices.OutputDeliveryModeUpload),
					URI:           to.Ptr("string"),
					Description:   to.Ptr("string"),
					JobOutputType: to.Ptr(armmachinelearningservices.JobOutputTypeURIFile),
				},
			},
			Settings: map[string]interface{}{},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/SweepJob/createOrUpdate.json
func ExampleJobsClient_CreateOrUpdate_createOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "string", armmachinelearningservices.JobBaseData{
		Properties: &armmachinelearningservices.SweepJob{
			Description: to.Ptr("string"),
			Properties: map[string]*string{
				"string": to.Ptr("string"),
			},
			Tags: map[string]*string{
				"string": to.Ptr("string"),
			},
			ComputeID:      to.Ptr("string"),
			DisplayName:    to.Ptr("string"),
			ExperimentName: to.Ptr("string"),
			JobType:        to.Ptr(armmachinelearningservices.JobTypeSweep),
			Services: map[string]*armmachinelearningservices.JobService{
				"string": &armmachinelearningservices.JobService{
					Endpoint:       to.Ptr("string"),
					JobServiceType: to.Ptr("string"),
					Port:           to.Ptr[int32](1),
					Properties: map[string]*string{
						"string": to.Ptr("string"),
					},
				},
			},
			EarlyTermination: &armmachinelearningservices.MedianStoppingPolicy{
				DelayEvaluation:    to.Ptr[int32](1),
				EvaluationInterval: to.Ptr[int32](1),
				PolicyType:         to.Ptr(armmachinelearningservices.EarlyTerminationPolicyTypeMedianStopping),
			},
			Limits: &armmachinelearningservices.SweepJobLimits{
				JobLimitsType:       to.Ptr(armmachinelearningservices.JobLimitsTypeSweep),
				MaxConcurrentTrials: to.Ptr[int32](1),
				MaxTotalTrials:      to.Ptr[int32](1),
				TrialTimeout:        to.Ptr("PT1S"),
			},
			Objective: &armmachinelearningservices.Objective{
				Goal:          to.Ptr(armmachinelearningservices.GoalMinimize),
				PrimaryMetric: to.Ptr("string"),
			},
			SamplingAlgorithm: &armmachinelearningservices.GridSamplingAlgorithm{
				SamplingAlgorithmType: to.Ptr(armmachinelearningservices.SamplingAlgorithmTypeGrid),
			},
			SearchSpace: map[string]interface{}{
				"string": map[string]interface{}{},
			},
			Trial: &armmachinelearningservices.TrialComponent{
				CodeID:  to.Ptr("string"),
				Command: to.Ptr("string"),
				Distribution: &armmachinelearningservices.Mpi{
					DistributionType:        to.Ptr(armmachinelearningservices.DistributionTypeMpi),
					ProcessCountPerInstance: to.Ptr[int32](1),
				},
				EnvironmentID: to.Ptr("string"),
				EnvironmentVariables: map[string]*string{
					"string": to.Ptr("string"),
				},
				Resources: &armmachinelearningservices.ResourceConfiguration{
					InstanceCount: to.Ptr[int32](1),
					InstanceType:  to.Ptr("string"),
					Properties: map[string]interface{}{
						"string": map[string]interface{}{
							"e6b6493e-7d5e-4db3-be1e-306ec641327e": nil,
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/cancel.json
func ExampleJobsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Cancel(ctx, "test-rg", "my-aml-workspace", "string", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
