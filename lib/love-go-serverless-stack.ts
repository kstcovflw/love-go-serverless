import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import { RestApi, LambdaIntegration } from 'aws-cdk-lib/aws-apigateway';

export class LoveGoServerlessStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // Hello Lambda function
    const helloFunction = new lambda.Function(this, 'HelloFunction', {
      code: lambda.Code.fromAsset('src/functions/hello'),
      handler: 'main',
      runtime: lambda.Runtime.PROVIDED_AL2023,
    });

    // The API Gateway
    const gateway = new RestApi(this, 'MyGateway', {
      defaultCorsPreflightOptions: {
        allowOrigins: ['*'],
        allowMethods: ['GET', 'POST'],
      },
    });

    // The Lambda integration
    const integration = new LambdaIntegration(helloFunction);

    // Creating the '/hello' resource
    const helloResource = gateway.root.addResource('hello');
    helloResource.addMethod('POST', integration); // POST method for the 'hello' resource
  }
}
