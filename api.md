# Functions

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#CreateFunctionUnionParam">CreateFunctionUnionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#EnrichConfigParam">EnrichConfigParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#EnrichStepParam">EnrichStepParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionType">FunctionType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#RouteListItemParam">RouteListItemParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#SplitFunctionSemanticPageItemClassParam">SplitFunctionSemanticPageItemClassParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#UpdateFunctionUnionParam">UpdateFunctionUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#EnrichConfig">EnrichConfig</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#EnrichStep">EnrichStep</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionUnion">FunctionUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionAudit">FunctionAudit</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionResponse">FunctionResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#ListFunctionsResponse">ListFunctionsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#RouteListItem">RouteListItem</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#SplitFunctionSemanticPageItemClass">SplitFunctionSemanticPageItemClass</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#UserActionSummary">UserActionSummary</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowUsageInfo">WorkflowUsageInfo</a>

Methods:

- <code title="post /v3/functions">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionNewParams">FunctionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/functions/{functionName}">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionName <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/functions/{functionName}">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pathFunctionName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionUpdateParams">FunctionUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/functions">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionListParams">FunctionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go/packages/pagination#FunctionsPage">FunctionsPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionUnion">FunctionUnion</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/functions/{functionName}">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionName <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Copy

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionCopyRequestParam">FunctionCopyRequestParam</a>

Methods:

- <code title="post /v3/functions/copy">client.Functions.Copy.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionCopyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionCopyNewParams">FunctionCopyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionVersionUnion">FunctionVersionUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#ListFunctionVersionsResponse">ListFunctionVersionsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionVersionGetResponse">FunctionVersionGetResponse</a>

Methods:

- <code title="get /v3/functions/{functionName}/versions/{versionNum}">client.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, versionNum <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionVersionGetParams">FunctionVersionGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionVersionGetResponse">FunctionVersionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/functions/{functionName}/versions">client.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionName <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#ListFunctionVersionsResponse">ListFunctionVersionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Calls

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#Call">Call</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#CallGetResponse">CallGetResponse</a>

Methods:

- <code title="get /v3/calls/{callID}">client.Calls.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#CallService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, callID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#CallGetResponse">CallGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/calls">client.Calls.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#CallService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#CallListParams">CallListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go/packages/pagination#CallsPage">CallsPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#Call">Call</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Errors

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#ErrorEvent">ErrorEvent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#InboundEmailEvent">InboundEmailEvent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#ErrorGetResponse">ErrorGetResponse</a>

Methods:

- <code title="get /v3/errors/{eventID}">client.Errors.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#ErrorService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#ErrorGetResponse">ErrorGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/errors">client.Errors.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#ErrorService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#ErrorListParams">ErrorListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go/packages/pagination#ErrorsPage">ErrorsPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#ErrorEvent">ErrorEvent</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Outputs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#AnyTypeUnion">AnyTypeUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#EventUnion">EventUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#OutputGetResponse">OutputGetResponse</a>

Methods:

- <code title="get /v3/outputs/{eventID}">client.Outputs.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#OutputService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#OutputGetResponse">OutputGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/outputs">client.Outputs.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#OutputService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#OutputListParams">OutputListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go/packages/pagination#OutputsPage">OutputsPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#EventUnion">EventUnion</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Workflows

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionVersionIdentifierParam">FunctionVersionIdentifierParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowRequestRelationshipParam">WorkflowRequestRelationshipParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#FunctionVersionIdentifier">FunctionVersionIdentifier</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#Workflow">Workflow</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowNewResponse">WorkflowNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowGetResponse">WorkflowGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowUpdateResponse">WorkflowUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowCopyResponse">WorkflowCopyResponse</a>

Methods:

- <code title="post /v3/workflows">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowNewParams">WorkflowNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowNewResponse">WorkflowNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/workflows/{workflowName}">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowName <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowGetResponse">WorkflowGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/workflows/{workflowName}">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowUpdateParams">WorkflowUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowUpdateResponse">WorkflowUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/workflows">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowListParams">WorkflowListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go/packages/pagination#WorkflowsPage">WorkflowsPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#Workflow">Workflow</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/workflows/{workflowName}">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowName <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v3/workflows/{workflowName}/call">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowService.Call">Call</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowCallParams">WorkflowCallParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#CallGetResponse">CallGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/workflows/copy">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowService.Copy">Copy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowCopyParams">WorkflowCopyParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowCopyResponse">WorkflowCopyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowVersionGetResponse">WorkflowVersionGetResponse</a>

Methods:

- <code title="get /v3/workflows/{workflowName}/versions/{versionNum}">client.Workflows.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, versionNum <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowVersionGetParams">WorkflowVersionGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowVersionGetResponse">WorkflowVersionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/workflows/{workflowName}/versions">client.Workflows.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#WorkflowVersionListParams">WorkflowVersionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go/packages/pagination#WorkflowVersionsPage">WorkflowVersionsPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go">bem</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bem-go#Workflow">Workflow</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
