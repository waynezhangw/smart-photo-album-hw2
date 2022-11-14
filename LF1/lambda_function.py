import json
import logging
import boto3
import time
import requests

logger = logging.getLogger()
logger.setLevel(logging.DEBUG)

def lambda_handler(event, context):
    # TODO implement
    logger.debug(event)
    logger.debug(context)
    logger.debug("upload photo successful!")
    for record in event['Records']:
        bucketName = record['s3']['bucket']['name']
        photoName = record['s3']['object']['key']
        # print("1111111")
        # print(bucketName, photoName)
        labels = []
        labels = get_photo_labels(bucketName, photoName)
        custom_labels = get_head_object_labels(bucketName, photoName)
        if len(custom_labels) > 0:
            custom_labels_list = custom_labels.split(",")
            print("custom_labels_list:", custom_labels_list)
            for i in range(len(custom_labels_list)):
                labels.append(custom_labels_list[i].strip())
            
        print("labels:", labels)
        new_doc = {
            "objectKey": photoName,
            "bucket": bucketName,
            "createdTimestamp": time.strftime("%Y%m%d-%H%M%S"),
            "labels": labels
        }
        index_into_es('photos','photo',new_doc)
    return {
        'statusCode': 200,
        'body': json.dumps('Hello from Lambda!')
    }
    
def get_photo_labels(bucketName, photoName):
    rekClient = boto3.client('rekognition')
    response = rekClient.detect_labels(Image={'S3Object':{'Bucket':bucketName,'Name':photoName}}, MaxLabels=10, MinConfidence=90)
    logger.debug("get_photo_labels response:")
    # logger.debug(response)
    # logger.debug(response['Labels'])
    labels = [label['Name'] for label in response['Labels']]
    logger.debug(labels)
    # ['Computer', 'Electronics', 'Tablet Computer']
    return labels

def get_head_object_labels(bucketName, photoName):
    s3Client = boto3.client('s3')
    head_object = s3Client.head_object(Bucket = bucketName, Key = photoName)
    logger.debug(head_object)
    custom_labels = []
    if 'customlabels' in head_object['Metadata']:
        custom_labels = head_object['Metadata']['customlabels']
    logger.debug(custom_labels)
    return custom_labels

def index_into_es(index, type_doc, new_doc):
    endpoint = 'https://search-photos-2mir7vk45ypl6bycqll62w3rjy.us-east-1.es.amazonaws.com/{}/{}'.format(index, type_doc)
    headers = {'Content-Type':'application/json'}
    # res = requests.post(endpoint, data=new_doc, headers=headers)
    res = requests.post(endpoint, auth = ('photos', 'Photos@123'), json = new_doc)
    logger.debug("index_into_es requests.post results:")
    logger.debug(res.content)