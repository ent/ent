/**
 * Copyright 2019-present Facebook Inc. All rights reserved.
 *
 * This source code is licensed under the Apache 2.0 license found
 * in the LICENSE file in the root directory of this source tree.
 *
 * @format
 */

const PropTypes = require('prop-types');
const React = require('react');

function SocialFooter(props) {
  const repoUrl = `https://github.com/${props.config.organizationName}/${props.config.projectName}`;
  return (
      <div className="footerSection">
        <h5>Social</h5>
        <div className="social">
          <a
              className="github-button"
              href={repoUrl}
              data-count-href={`${repoUrl}/stargazers`}
              data-show-count="true"
              data-count-aria-label="# stargazers on GitHub"
              aria-label="Star this project on GitHub">
            {props.config.projectName}
          </a>
        </div>
        {props.config.twitterUsername && (
            <div className="social">
              <a
                  href={`https://twitter.com/${props.config.twitterUsername}`}
                  className="twitter-follow-button">
                Follow @{props.config.twitterUsername}
              </a>
            </div>
        )}
        {props.config.facebookAppId && (
            <div className="social">
              <div
                  className="fb-like"
                  data-href={props.config.url}
                  data-colorscheme="dark"
                  data-layout="standard"
                  data-share="true"
                  data-width="225"
                  data-show-faces="false"
              />
            </div>
        )}
      </div>
  );
}

SocialFooter.propTypes = {
  config: PropTypes.object,
};

class Footer extends React.Component {
  render() {
    const docsPart = `${
        this.props.config.docsUrl ? `${this.props.config.docsUrl}/` : ''
    }`;
    return (
        <footer className="nav-footer" id="footer">
          <section className="sitemap">
            {this.props.config.footerIcon && (
                <a href={this.props.config.baseUrl} className="nav-home">
                  <img
                      src={`${this.props.config.baseUrl}${this.props.config.footerIcon}`}
                      alt={this.props.config.title}
                      width="66"
                      height="58"
                  />
                </a>
            )}
            <div className="footerSection">
              <h5>Docs</h5>
              <a
                  href={`
                ${this.props.config.baseUrl}${docsPart}${this.props.language}/getting-started`}>
                Getting Started
              </a>
              <a
                  href={`
                ${this.props.config.baseUrl}${docsPart}${this.props.language}/schema-def`}>
                Schema Guide
              </a>
              <a
                  href={`
                ${this.props.config.baseUrl}${docsPart}${this.props.language}/code-gen`}>
                Code Generation
              </a>
              <a
                  href={`
                ${this.props.config.baseUrl}${docsPart}${this.props.language}/graphql`}>
                GraphQL Integration
              </a>
              <a
                  href={`
                ${this.props.config.baseUrl}${docsPart}${this.props.language}/migrate`}>
                Schema Migration
              </a>
            </div>
            <div className="footerSection">
              <h5>Community</h5>
              <a href={`${this.props.config.githubRepo}`} target="_blank">
                GitHub
              </a>
              <a href={`${this.props.config.slackChannel}`} target="_blank">
                Slack
              </a>
              <a href={`${this.props.config.newsletter}`} target="_blank">
                Newsletter
              </a>
              <a href={`${this.props.config.githubRepo}/discussions`} target="_blank">
                Discussions
              </a>
            </div>
            <div className="footerSection">
              <h5>Legal</h5>
              <a
                  href="https://opensource.facebook.com/legal/privacy/"
                  target="_blank"
                  rel="noreferrer noopener">
                Privacy
              </a>
              <a
                  href="https://opensource.facebook.com/legal/terms/"
                  target="_blank"
                  rel="noreferrer noopener">
                Terms
              </a>
              <a
                  href="https://opensource.facebook.com/legal/data-policy/"
                  target="_blank"
                  rel="noreferrer noopener">
                Data Policy
              </a>
              <a
                  href="https://opensource.facebook.com/legal/cookie-policy/"
                  target="_blank"
                  rel="noreferrer noopener">
                Cookie Policy
              </a>
            </div>
            <SocialFooter config={this.props.config} />
          </section>
          <a
              href="https://opensource.facebook.com/"
              target="_blank"
              rel="noreferrer noopener"
              className="fbOpenSource">
            <img
                src={`${this.props.config.baseUrl}img/oss_logo.png`}
                alt="Facebook Open Source"
                width="170"
                height="45"
            />
          </a>
          <section className="copyright">
            {this.props.config.copyright && (
                <span>{this.props.config.copyright}</span>
            )}{' '}
            <br/>
            <br/>
            The Go gopher was designed by <a href="http://reneefrench.blogspot.com/" style={{display: 'inline'}}> Renee French </a>.
            The design is licensed under the Creative Commons 3.0 Attributions license. Read this{' '}
            <a href="https://blog.golang.org/gopher" style={{display: 'inline'}}> article </a>{' '} for more details.
            <br/>
            Design by Moriah Rich, illustration by Ariel Mashraki.
          </section>
        </footer>
    );
  }
}

module.exports = Footer;