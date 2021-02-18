/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @format
 */

const React = require('react');

const CompLibrary = require('../../core/CompLibrary.js');

const MarkdownBlock = CompLibrary.MarkdownBlock; /* Used to read markdown */
const Container = CompLibrary.Container;
const GridBlock = CompLibrary.GridBlock;

const arrow = '\u2192';

const Block = props => (
  <div className="block">
    <div className="blockTitle">
      <a href={props.link}>
        <div className="blockTitleText">{props.title}</div>{' '}
      </a>
      <a className="yellowArrow" href={props.link}>{arrow}</a>
    </div>
    <div className="blockContent">{props.content}</div>
  </div>
);

const Features = () => (
  <div className="features">
    <Block
      title="Schema As Code"
      content="Simple API for modeling any database schema as Go objects"
      link="docs/schema-def"
    />
    <Block
      title="Easily Traverse Any Graph"
      content="Run queries, aggregations and traverse any graph structure easily"
      link="docs/traversals"
    />
    <Block
      title="Statically Typed And Explicit API"
      content="100% statically typed and explicit API using code generation"
      link="docs/code-gen"
    />
  </div>
);

class HomeSplash extends React.Component {
  render() {
    const {siteConfig, language = ''} = this.props;
    const {baseUrl, docsUrl} = siteConfig;
    const docsPart = `${docsUrl ? `${docsUrl}/` : ''}`;
    const langPart = `${language ? `${language}/` : ''}`;
    const docUrl = doc => `${baseUrl}${docsPart}${langPart}${doc}`;

    const SplashContainer = props => (
      <div className="homeContainer">
        <div className="homeSplashFade">
          <div className="wrapper homeWrapper">{props.children}</div>
        </div>
      </div>
    );

    const Logo = props => (
      <div className="projectLogo">
        <img src={props.img_src} alt="Project Logo" />
      </div>
    );

    const ProjectTitle = () => (
      <div>
        <div className="projectTitleContainer">
          <img src="https://entgo.io/assets/logo.png" />
          <div className="projectTitle">
            <p>{siteConfig.tagline}</p>
          </div>
        </div>
        <p className="projectDesc">
          Simple, yet powerful ORM for modeling and querying data.
        </p>
      </div>
    );

    const PromoSection = props => (
      <div className="section promoSection">
        <div className="promoRow">
          <div className="pluginRowBlock">{props.children}</div>
        </div>
      </div>
    );

    const Button = props => (
      <div className="pluginWrapper buttonWrapper">
        <a className="button" href={props.href} target={props.target}>
          {props.children}
        </a>
      </div>
    );

    return (
      <SplashContainer>
        {/*<Logo img_src={`${baseUrl}img/undraw_monitor.svg`} />*/}
        <div className="inner">
          <ProjectTitle siteConfig={siteConfig} />
          <div className="gettingStartedButton">
            <a href="./docs/getting-started">
              <div className="gettingStartedButtonText">
                <div className="gettingStartedText">{'Getting Started '}</div>
                <div className="gettingStartedButtonArrow">{arrow}</div>
              </div>
            </a>
          </div>
          <div className="gopherGraph">
            <img src="https://entgo.io/assets/gopher_graph.png" />
          </div>
          <Features />
        </div>
      </SplashContainer>
    );
  }
}

class Index extends React.Component {
  render() {
    const {config: siteConfig, language = ''} = this.props;
    const {baseUrl} = siteConfig;

    const Showcase = () => {
      if ((siteConfig.users || []).length === 0) {
        return null;
      }

      const showcase = siteConfig.users
        .filter(user => user.pinned)
        .map(user => (
          <a href={user.infoLink} key={user.infoLink}>
            <img src={user.image} alt={user.caption} title={user.caption} />
          </a>
        ));

      const pageUrl = page => baseUrl + (language ? `${language}/` : '') + page;

      return (
        <div className="productShowcaseSection paddingBottom">
          <h2>Who is Using This?</h2>
          <p>This project is used by all these people</p>
          <div className="logos">{showcase}</div>
          <div className="more-users">
            <a className="button" href={pageUrl('users.html')}>
              More {siteConfig.title} Users
            </a>
          </div>
        </div>
      );
    };

    return (
      <div>
        <HomeSplash siteConfig={siteConfig} language={language} />
      </div>
    );
  }
}

module.exports = Index;
