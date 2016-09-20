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
 * @param Bool recursive
 * 
 * @returns String|Error 
 */
import InvalidLinkError from './ErrorClass.js';

export function ValidateUrl(input, recursive = false){
        let parser = document.createElement('a');
        parser.href = input;
        
        const puncRegex = /[,.?\-]/;
        const relativeHostname = location.hostname;
        
        let href = parser.href;
        let proto = parser.protocol;
        let hostname = parser.hostname;
        let returnedURL;
        
        if(recursive && relativeHostname === hostname){
                throw new InvalidLinkError("Don't Shorten my urls");
        }
        
        /* Make sure there is a valid protocol and it's not too long' */
        if(proto.length > 0 && hostname.length > 0 && hostname.length < (126)
          && href.length < 2083){
                /* Make sure url is absolute and not relative */
                if(relativeHostname === hostname){
                        /* Add http and recursivly check if valid */
                        let absHostname = `${href.replace(relativeHostname + '/', '')}`;
                        return ValidateUrl(absHostname, true);
                }
                
                /* Hostnames cant start with a punctuation */
                if(puncRegex.test(hostname.charAt(0))){
                        throw new InvalidLinkError("Links cant start with punctuation");
                }
                
                /* But they must have a period somewhere */
                if(hostname.indexOf('.') === -1){
                        throw new InvalidLinkError("The link contains an error");
                }
        }
        else{
                throw new InvalidLinkError("Link is too large");
        }
        
        returnedURL = href;
        parser = null;
        return returnedURL;
}