using Pulumi;
using Pulumi.AzureAD;
using Pulumi.Random;
using System.Threading.Tasks;

class MyStack : Stack
{
    public MyStack()
    {
        var serverRandomPet = new RandomPet("random-name");

        var randomString = new RandomString("random", new RandomStringArgs
        {
            Length = 6,
            Special = false,
        });

        var user = new User("csharp", new UserArgs
        {
            DisplayName = serverRandomPet.Id,
            MailNickname = randomString.Result,
            Password = "SecretP@sswd99!",
            UserPrincipalName = Output.Format($"{randomString.Result}@pulumici.onmicrosoft.com"),
        });

        this.UserId = user.Id;
    }

    [Output]
    public Output<string> UserId { get; set; }
}

class Program
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}