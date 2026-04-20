# Functions

Params Types:

- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#ClassificationListItemParam">ClassificationListItemParam</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#CreateFunctionUnionParam">CreateFunctionUnionParam</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#EnrichConfigParam">EnrichConfigParam</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#EnrichStepParam">EnrichStepParam</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionType">FunctionType</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#SplitFunctionSemanticPageItemClassParam">SplitFunctionSemanticPageItemClassParam</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#UpdateFunctionUnionParam">UpdateFunctionUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#ClassificationListItem">ClassificationListItem</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#EnrichConfig">EnrichConfig</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#EnrichStep">EnrichStep</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionUnion">FunctionUnion</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionAudit">FunctionAudit</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionResponse">FunctionResponse</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#ListFunctionsResponse">ListFunctionsResponse</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#SplitFunctionSemanticPageItemClass">SplitFunctionSemanticPageItemClass</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#UserActionSummary">UserActionSummary</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowUsageInfo">WorkflowUsageInfo</a>

Methods:

- <code title="post /v3/functions">client.Functions.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionNewParams">FunctionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/functions/{functionName}">client.Functions.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionName <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/functions/{functionName}">client.Functions.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pathFunctionName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionUpdateParams">FunctionUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/functions">client.Functions.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionListParams">FunctionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk/packages/pagination#FunctionsPage">FunctionsPage</a>[<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionUnion">FunctionUnion</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/functions/{functionName}">client.Functions.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionName <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Copy

Params Types:

- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionCopyRequestParam">FunctionCopyRequestParam</a>

Methods:

- <code title="post /v3/functions/copy">client.Functions.Copy.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionCopyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionCopyNewParams">FunctionCopyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionVersionUnion">FunctionVersionUnion</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#ListFunctionVersionsResponse">ListFunctionVersionsResponse</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionVersionGetResponse">FunctionVersionGetResponse</a>

Methods:

- <code title="get /v3/functions/{functionName}/versions/{versionNum}">client.Functions.Versions.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, versionNum <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionVersionGetParams">FunctionVersionGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionVersionGetResponse">FunctionVersionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/functions/{functionName}/versions">client.Functions.Versions.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionName <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#ListFunctionVersionsResponse">ListFunctionVersionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Calls

Response Types:

- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#Call">Call</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#CallGetResponse">CallGetResponse</a>

Methods:

- <code title="get /v3/calls/{callID}">client.Calls.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#CallService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, callID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#CallGetResponse">CallGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/calls">client.Calls.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#CallService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#CallListParams">CallListParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk/packages/pagination#CallsPage">CallsPage</a>[<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#Call">Call</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Errors

Response Types:

- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#ErrorEvent">ErrorEvent</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#InboundEmailEvent">InboundEmailEvent</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#ErrorGetResponse">ErrorGetResponse</a>

Methods:

- <code title="get /v3/errors/{eventID}">client.Errors.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#ErrorService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#ErrorGetResponse">ErrorGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/errors">client.Errors.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#ErrorService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#ErrorListParams">ErrorListParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk/packages/pagination#ErrorsPage">ErrorsPage</a>[<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#ErrorEvent">ErrorEvent</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Outputs

Response Types:

- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#AnyTypeUnion">AnyTypeUnion</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#EventUnion">EventUnion</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#OutputGetResponse">OutputGetResponse</a>

Methods:

- <code title="get /v3/outputs/{eventID}">client.Outputs.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#OutputService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#OutputGetResponse">OutputGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/outputs">client.Outputs.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#OutputService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#OutputListParams">OutputListParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk/packages/pagination#OutputsPage">OutputsPage</a>[<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#EventUnion">EventUnion</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Workflows

Params Types:

- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionVersionIdentifierParam">FunctionVersionIdentifierParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#FunctionVersionIdentifier">FunctionVersionIdentifier</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#Workflow">Workflow</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowAudit">WorkflowAudit</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowEdgeResponse">WorkflowEdgeResponse</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowNodeResponse">WorkflowNodeResponse</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowNewResponse">WorkflowNewResponse</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowGetResponse">WorkflowGetResponse</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowUpdateResponse">WorkflowUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowCopyResponse">WorkflowCopyResponse</a>

Methods:

- <code title="post /v3/workflows">client.Workflows.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowNewParams">WorkflowNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowNewResponse">WorkflowNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/workflows/{workflowName}">client.Workflows.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowName <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowGetResponse">WorkflowGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/workflows/{workflowName}">client.Workflows.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowUpdateParams">WorkflowUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowUpdateResponse">WorkflowUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/workflows">client.Workflows.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowListParams">WorkflowListParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk/packages/pagination#WorkflowsPage">WorkflowsPage</a>[<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#Workflow">Workflow</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/workflows/{workflowName}">client.Workflows.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowName <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v3/workflows/{workflowName}/call">client.Workflows.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowService.Call">Call</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowCallParams">WorkflowCallParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#CallGetResponse">CallGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/workflows/copy">client.Workflows.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowService.Copy">Copy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowCopyParams">WorkflowCopyParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowCopyResponse">WorkflowCopyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowVersionGetResponse">WorkflowVersionGetResponse</a>

Methods:

- <code title="get /v3/workflows/{workflowName}/versions/{versionNum}">client.Workflows.Versions.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, versionNum <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowVersionGetParams">WorkflowVersionGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowVersionGetResponse">WorkflowVersionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/workflows/{workflowName}/versions">client.Workflows.Versions.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#WorkflowVersionListParams">WorkflowVersionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk/packages/pagination#WorkflowVersionsPage">WorkflowVersionsPage</a>[<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#Workflow">Workflow</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InferSchema

Response Types:

- <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#InferSchemaNewResponse">InferSchemaNewResponse</a>

Methods:

- <code title="post /v3/infer-schema">client.InferSchema.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#InferSchemaService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#InferSchemaNewParams">InferSchemaNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk">bem</a>.<a href="https://pkg.go.dev/github.com/bem-team/bem-go-sdk#InferSchemaNewResponse">InferSchemaNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
