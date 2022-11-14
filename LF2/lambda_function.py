import json
import boto3
import time
import requests
import inflect
import logging

logger = logging.getLogger()
logger.setLevel(logging.DEBUG)

def lambda_handler(event, context):
    # TODO implement
    print(event)
    logger.debug("input text is:")
    logger.debug(event)
    inputText = event['queryStringParameters']['q']
    # inputText = event['q']
    keywords = get_keywords(inputText)
    keywords = convert_keywords_2_plural(keywords)
    # ['https://wz2055-cloud-hw2.s3.amazonaws.com/pexels-gilberto-reyes-825947.jpg', 'https://wz2055-cloud-hw2.s3.amazonaws.com/pexels-pixabay-416160.jpg', 'https://wz2055-cloud-hw2.s3.amazonaws.com/pexels-pixabay-45201.jpg', 'https://wz2055-cloud-hw2.s3.amazonaws.com/pexels-pixabay-220938.jpg', 'https://wz2055-cloud-hw2.s3.amazonaws.com/pexels-valeria-boltneva-1805164.jpg', 'https://wz2055-cloud-hw2.s3.amazonaws.com/pexels-evg-kowalievska-1170986.jpg', 'https://wz2055-cloud-hw2.s3.amazonaws.com/pexels-pixabay-104827.jpg', 'https://wz2055-cloud-hw2.s3.amazonaws.com/pexels-chevanon-photography-1108099.jpg', 'https://wz2055-cloud-hw2.s3.amazonaws.com/pexels-muhannad-alatawi-58997.jpg']
    image_array = get_image_locations(keywords)
    print(image_array)
    return {
        'statusCode': 200,
        'headers':{
            'Access-Control-Allow-Origin':'*',
            'Access-Control-Allow-Credentials':True
        },
        'body': json.dumps({"results":image_array})
    }

def get_keywords(inputText):
    lex = boto3.client('lex-runtime')
    response = lex.post_text(
        botName = 'photoSearchBot',
        botAlias = 'abcd',
        userId = 'wayne',
        inputText = inputText
    )
    logger.debug("cat-------------------------")
    logger.debug(response)
    # print(response['slots'])
    keywords = []
    # this check is very important, since sometimes lex cannot recognize slot input
    if "slots" not in response:
        return [inputText]
    slots = response['slots']
    keywords = [v for _, v in slots.items() if v]
    print(keywords)
    return keywords

def convert_keywords_2_plural(keywords):
    p = inflect.engine()
    new_keywords = set(keywords)
    for word in keywords:
        if p.plural(word) == False:
            new_keywords.add(word)
        else:
            new_keywords.add(p.plural(word))
        
        if p.singular_noun(word) == False:
            new_keywords.add(word)
        else:
            new_keywords.add(p.singular_noun(word))
    
    print(new_keywords)
    return list(new_keywords)

def get_image_locations(keywords):
    baseUrl = 'https://search-photos-2mir7vk45ypl6bycqll62w3rjy.us-east-1.es.amazonaws.com'
    path = '/photos/photo/_search'
    url = baseUrl + path
    auth = ('photos', 'Photos@123')
    
    headers = {'Content-Type': 'application/json'}
    prepared_q = []
    for k in keywords:
        prepared_q.append({"match": {"labels": k}})
    q = {"query": {"bool": {"should": prepared_q}}}
    response = requests.get(url, auth=auth, headers=headers, data=json.dumps(q))
    print(response)
    print(response.json())
    
    image_array = []
    for each in response.json()['hits']['hits']:
        objectKey = each['_source']['objectKey']
        bucket = each['_source']['bucket']
        image_url = "https://" + bucket + ".s3.amazonaws.com/" + objectKey
        image_array.append(image_url)
        print(each['_source']['labels'])
    print(image_array)
    image_array = set(image_array)
    return list(image_array)