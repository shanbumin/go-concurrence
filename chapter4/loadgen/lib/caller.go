package lib

import (
	"time"
)

//最后，应该让载荷发生器的使用者可以根据具体需求对它进行适当的扩展和定制。为此，需要预先在其结构中添加一个字段 caller  lib.Caller，并以此作为载荷发生器的扩展接口。
//不过，在添加这个字段之前，应该搞清楚主次、搞清来龙去脉。首先，载荷发生器的核心功能，肯定是控制和协调载荷的生成和发送、响应的接收和验证，以及最终结果的递交等一系列操作。
//核心功能必然是大流程上的控制，那么这些流程中涉及到的一些具体的操作、功能是否要组件化呢？这里我们就一起看看有哪些可以作为组件功能。
//显然，我不知道或者无法预测被测软件提供API的形式。况且，载荷发生器不应该对此有所约束，它们可以是任意的。
//因此，与调用被测软件API有关的功能应该作为组件功能，这涉及请求的发送操作和响应的接收操作。并且，既然要组件化调用被测软件API的功能，那么请求的生成操作和响应的检查操作，也肯定无法由载荷发生器本身来提供。



//根据上面的分析，我编写了这样一个接口类型来体现可组件化的功能
//todo 虽然被视为非核心功能，但是该接口类型中的那几个方法所代表的操作，也都是载荷发生器在运行过程中不可或缺的。
//todo 因此我们应该确保在初始化载荷发生器时，持有一个lib.Caller接口类型的实现值。在载荷发生器的结构中存在一个该类型的字段，以便存放这个实现值。

// Caller 表示调用器的接口。
type Caller interface {
	BuildReq() RawReq // 构建请求。
	Call(req []byte, timeoutNS time.Duration) ([]byte, error) // 调用。
	CheckResp(rawReq RawReq, rawResp RawResp) *CallResult  // 检查响应。
}


