package main

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/mohan2020coder/go-workflows/tester"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Workflow(t *testing.T) {
	tester := tester.NewWorkflowTester[string](Workflow1)

	tester.OnActivity(Activity1, mock.Anything, 35, 12).Return(47, nil)

	tester.Execute(context.Background(), "Hello world"+uuid.NewString())

	require.True(t, tester.WorkflowFinished())

	wr, werr := tester.WorkflowResult()
	require.Equal(t, "result", wr)
	require.NoError(t, werr)
	tester.AssertExpectations(t)
}
