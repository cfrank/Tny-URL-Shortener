import * as Constants from './constants.js';
import {InvalidLinkError, NoticeUserError, GeneralError} from './ErrorClass.js';

/*
 * Parse and validate the url
 * 
 * This will check the url against some quite lenient rules
 * and try to only allow valid urls to pass but also allow people
 * to send through whatever crazy url they want to create.
 * 
 * All urls are returned URIEncoded
 * 
 * @param String input
 * 
 * @returns String|Error
 */
export function ValidateUrl(input){
        let parser = document.createElement('a');
        parser.href = input;
        
        const puncRegex = /[,.?\-*$%^&()_=+]/;
        const relativeHostname = location.hostname;
        const defaultProto = 'http://';
        
        // Make sure url is absolute
        if(parser.hostname === relativeHostname)
                parser.href = `${defaultProto}${input}`;
        
        let href = parser.href;
        let proto = parser.protocol;
        let hostname = parser.hostname;
        let tld = hostname.split('.').pop();
        
        /* Make sure there is a valid protocol and it's not too long' */
        if(proto.length > 0 && hostname.length > 0 && hostname.length < (126)
          && href.length < 2083){
                /* Hostnames cant start with a punctuation */
                if(puncRegex.test(hostname.replace('www.','').charAt(0))){
                        throw new InvalidLinkError("Links cant start with punctuation");
                }
                
                /* But they must have a period somewhere */
                if(hostname.indexOf('.') === -1){
                        throw new InvalidLinkError("The link contains an error");
                }
                
                /* TLD checks */
                if(tld === null || tld.length < 2 || tld.length > 64 || puncRegex.test(tld)){
                        throw new InvalidLinkError("The link has an invalid TLD");
                }
        }
        else{
                throw new InvalidLinkError("Link is of incorrect length");
        }
        
        let returnedURL = href;
        parser = null;
        return returnedURL;
}

/*
 * Validate a linkid
 * 
 * Make sure the linkid length is correct.
 * TODO: add more checks
 */
export function ValidateLinkid(linkid){
        if(linkid.length === 6){
                // Correct length
                return linkid;
        }
        else{
                throw new GeneralError("Linkid is of incorrect length");
        }
}

/*
 * Fetch json data from specified href argument
 * 
 * @returns Promise|Error
 */
export function CallFetchJson(url){
        return fetch(url).then(function(response){
                return response.json();
        }).catch(function(error){
                const fetchError = new NoticeUserError("ERROR: Could not fetch data from server! Please Reload", false);
                fetchError.show();
                return false;
        });
}

/*
 * Send json as POST request to provided url
 */
export function JsonPostRequest(url, data){
        return fetch(url,
        {
                headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json'
                },
                method: 'POST',
                body: JSON.stringify(data)
        }).then(function(response){
                return response.json();
        }).catch(function(error){
                const postError = new NoticeUserError("ERROR: Can't connect to the server! Please Reload", false);
                postError.show();
                return false;
        });
}