/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @format
 */

const React = require('react');

class Footer extends React.Component {
  docUrl(doc, language) {
    const baseUrl = this.props.config.baseUrl;
    const docsUrl = this.props.config.docsUrl;
    const docsPart = `${docsUrl ? `${docsUrl}/` : ''}`;
    const langPart = `${language ? `${language}/` : ''}`;
    return `${baseUrl}${docsPart}${langPart}${doc}`;
  }

  pageUrl(doc, language) {
    const baseUrl = this.props.config.baseUrl;
    return baseUrl + (language ? `${language}/` : '') + doc;
  }

  render() {
    return (
      <footer className="nav-footer" id="footer">
        <section className="sitemap">
          <div>
            <h5>Docs</h5>
            <a href={this.docUrl('getting-started')}>
              Getting Started
            </a>
            <a href={this.docUrl('schema-def')}>
              Schema Guide
            </a>
            <a href={this.docUrl('code-gen')}>
              Code Generation
            </a>
            <a href={this.docUrl('migrate')}>
             Schema Migration
            </a>
            <br/>
            <h5>Legal</h5>
            <a href="https://opensource.facebook.com/legal/privacy/" target="_blank" rel="noreferrer noopener">Privacy</a>
            <a href="https://opensource.facebook.com/legal/terms/" target="_blank" rel="noreferrer noopener">Terms</a>
          </div>
          <div>
            <h5>Credit</h5>
            <span className="copyright">
              The Go gopher was designed by{' '}
              <a
                href="http://reneefrench.blogspot.com/"
                style={{display: 'inline'}}>
                Renee French
              </a>
              . The design is licensed under the Creative Commons 3.0
              Attributions license. Read this{' '}
              <a
                href="https://blog.golang.org/gopher"
                style={{display: 'inline'}}>
                article
              </a>{' '}
              for more details.
            </span>
            <br/><br/>
            <span className="copyright">
              Design by Moriah Rich, illustration by Ariel Mashraki.
            </span>
          </div>
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
        <section className="copyright">{this.props.config.copyright}</section>
      </footer>
    );
  }
}

module.exports = Footer;
